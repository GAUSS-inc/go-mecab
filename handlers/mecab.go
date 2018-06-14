// Copyright 2018 GAUSS All Rights Reserved.
// 形態素解析

package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"gauss/go-mecab/application"
	"gauss/go-mecab/dto/request"
	"gauss/go-mecab/dto/response"

	"github.com/fatih/structs"
	"github.com/labstack/echo"
	"github.com/shogo82148/go-mecab"
)

// MecabRouter ルーター設定
func MecabRouter(e *echo.Echo) {
	g := e.Group("/mecab")

	g.GET("", application.AppHandler(getParse))
}

// getParse godoc
// @Summary 形態素解析
// @Description 形態素解析結果を取得する。
// @Param sentence query string true "形態素解析対象"
// @Success 200 {object} response.MecabResults
// @Router /mecab [get]
func getParse(c *application.AppContext) error {
	param := new(request.MecabParam)
	if err := c.BindValidate(param); err != nil {
		return err
	}
	tagger, err := mecab.New(map[string]string{"output-format-type": "wakati"})
	defer tagger.Destroy()

	node, err := tagger.ParseToNode(param.Sentence)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error_message": err.Error()})
	}

	res := &response.MecabResults{Results: make([]*response.MecabResult, 0)}
	for ; node != (mecab.Node{}); node = node.Next() {
		features := strings.Split(node.Feature(), ",")
		if node.Surface() == "" || features[0] == "BOS/EOS" {
			continue
		}
		elements := &response.MecabResult{Surface: node.Surface()}
		es := structs.New(elements)
		for idx, feature := range features {
			for _, fld := range es.Fields() {
				fno, err := strconv.Atoi(fld.Tag("feature"))
				if err != nil {
					continue
				}
				if idx == fno {
					fld.Set(feature)
					break
				}
			}
		}

		res.Results = append(res.Results, elements)
	}

	return c.JSON(http.StatusOK, res)
}

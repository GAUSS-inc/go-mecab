// Copyright 2018 GAUSS All Rights Reserved.
// echo.Context ラッパー

package application

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// AppContext echo.Context をラップ
type AppContext struct {
	echo.Context
}

// BindValidate BindとValidateを合わせたメソッド
func (c *AppContext) BindValidate(i interface{}) error {
	if err := c.Bind(i); err != nil {
		return c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		// return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Validate(i); err != nil {
		return c.String(http.StatusBadRequest, "Validate is failed: "+err.Error())
		// return c.NoContent(http.StatusBadRequest)
	}
	return nil
}

// callFunc
type callFunc func(c *AppContext) error

// AppHandler Router登録用
func AppHandler(h callFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println(fmt.Sprintf("debug: start: %v %v", c.Request().Method, c.Request().RequestURI))
		return h(c.(*AppContext))
	}
}

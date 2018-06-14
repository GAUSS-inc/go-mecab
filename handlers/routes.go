// Copyright 2018 GAUSS All Rights Reserved.
// ルーティング設定

package handlers

import (
	_ "gauss/go-mecab/docs"

	"github.com/labstack/echo"
	"github.com/swaggo/echo-swagger"
)

// Router ルーティング設定
func Router(e *echo.Echo) {
	MecabRouter(e)

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

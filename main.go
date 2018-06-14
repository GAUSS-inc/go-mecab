// Copyright 2018 GAUSS All Rights Reserved.
// API サーバメイン
// @title MeCab API
// @version 1.0.0
// @description Golang MeCab APIサーバ
// @contact.name GAUSS
// @contact.url https://gauss-ai.jp
// @contact.email support@gauss-ai.jp
// @host localhost:1323
// @BasePath /

package main

import (
	"log"

	"gauss/go-mecab/application"
	"gauss/go-mecab/handlers"

	"github.com/comail/colog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	//===================================
	// Initialize
	//===================================
	colog.SetFormatter(&colog.StdFormatter{
		Flag: log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()
	e := echo.New()
	e.Validator = application.GetValidator()

	// Router初期化
	handlers.Router(e)

	//===================================
	// Middleware
	//===================================
	// echo.Context をラップして扱うために middleware として登録する
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return h(&application.AppContext{c})
		}
	})
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${host} [${time_rfc3339}] \"${method} ${uri}\" ${status} ${bytes_in} ${bytes_out}\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	//===================================
	// server start
	//===================================
	e.Start(":1323")
}

package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() // 得到一个 echo.Echo 的实例

	// 注册路由
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// 开启 HTTP Server
	e.Logger.Fatal(e.Start(":4789"))
}

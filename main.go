package main

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New() // 得到一个 echo.Echo 的实例
	// e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// 注册路由
	e.GET("/", func(c echo.Context) error {
		_signature := c.QueryParam("signature")
		_timestamp := c.QueryParam("timestamp")
		_nonce := c.QueryParam("nonce")
		_token := "fengfenglovejiangjiang"
		_tmpArr := []string{_timestamp, _nonce, _token}
		sort.Strings(_tmpArr)
		newstr := strings.Join(_tmpArr, "")
		a := sha1.New()
		a.Write([]byte(newstr))
		if hex.EncodeToString(a.Sum(nil)) == _signature {
			return c.String(http.StatusOK, "Hello, World!")
		} else {
			return c.String(http.StatusMultipleChoices, "aaa")
		}
	})

	// 开启 HTTP Server
	e.Logger.Fatal(e.Start(":8000"))
}

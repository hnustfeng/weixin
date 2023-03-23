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
		_echostr := c.QueryParam("echostr")
		_token := "fengfenglovejiangjiang"
		_tmpArr := []string{_timestamp, _nonce, _token}
		sort.Strings(_tmpArr)
		newstr := strings.Join(_tmpArr, "")
		a := sha1.New()
		n := len(_timestamp) + len(_nonce) + len(_token)
		var b strings.Builder
		b.Grow(n)
		for i := 0; i < len(_tmpArr); i++ {
			b.WriteString(_tmpArr[i])
		}
		a.Write([]byte(newstr))
		if hex.EncodeToString(a.Sum(nil)) == _signature {
			return c.String(http.StatusOK, _echostr)
		} else {
			return c.String(500, "")
		}
	})

	// 开启 HTTP Server
	e.Logger.Fatal(e.Start(":80"))
}

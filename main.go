package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
	"weixin/service"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default() // 得到一个 echo.Echo 的实例
	// e.Use(middleware.CORS())

	// 注册路由
	e.GET("/", func(c *gin.Context) {
		_signature := c.Query("signature")
		_timestamp := c.Query("timestamp")
		_nonce := c.Query("nonce")
		_echostr := c.Query("echostr")
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
			c.String(http.StatusOK, _echostr)
		} else {
			c.String(500, "")
		}
	})

	e.POST("/", WXMsgReceive)
	e.GET("/weather", func(c *gin.Context) {
		data, err := service.GetAir()
		if err != nil {
			c.String(500, err.Error())
		}
		c.JSON(http.StatusOK, data)
	})

	e.GET("/time", func(c *gin.Context) {
		timenow := time.Now()
		time, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-12-02 00:00:00", time.Local)
		day := int(timenow.Sub(time).Hours() / 24)
		data := make(map[string]interface{})
		data["day"] = day
		fmt.Println(time)
		c.JSON(200, data)
	})
	// 开启 HTTP Server
	e.Run(":8000")
}

// WXTextMsg 微信文本消息结构体
type WXTextMsg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        int64  `xml:"MsgId"`
}

// WXMsgReceive 微信消息接收
func WXMsgReceive(c *gin.Context) {
	var textMsg WXTextMsg
	err := c.ShouldBindXML(&textMsg)
	if err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		return
	}

	log.Printf("[消息接收] - 收到消息, 消息类型为: %s, 消息内容为: %s\n", textMsg.MsgType, textMsg.Content)
}

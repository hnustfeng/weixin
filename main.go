package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
	"weixin/service"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func main() {
	spec := "0 0 8 * * *" // 每天早晨8:00
	c := cron.New()
	c.AddFunc(spec, service.Send)
	c.Start()
	fmt.Println("开启定时任务")
	select {}
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
		data, err := service.GetIndices()
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
	//weather()
	//everydaysen()
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

func WXMsgReceive(c *gin.Context) {
	var textMsg WXTextMsg
	err := c.ShouldBindXML(&textMsg)
	if err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		return
	}

	log.Printf("[消息接收] - 收到消息, 消息类型为: %s, 消息内容为: %s\n", textMsg.MsgType, textMsg.Content)

	// 对接收的消息进行被动回复
	WXMsgReply(c, textMsg.ToUserName, textMsg.FromUserName)
}

// WXRepTextMsg 微信回复文本消息结构体
type WXRepTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}

// WXMsgReply 微信消息回复
func WXMsgReply(c *gin.Context, fromUser, toUser string) {
	repTextMsg := WXRepTextMsg{
		ToUserName:   toUser,
		FromUserName: fromUser,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      fmt.Sprintf("[消息回复] - %s", time.Now().Format("2006-01-02 15:04:05")),
	}

	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
}

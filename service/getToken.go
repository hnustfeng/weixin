package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

var (
	APPID          = "wxde0cca8d448d5fba"
	APPSECRET      = "4e1c749696b8e23d22948b9f16cd88e7"
	SentTemplateID = "UEn1u34JHcfSZnGQbtyGdn4Lr7EVJWSl_6hP18jSBvA"
)

type token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

//获取微信accesstoken
func GetAccessToken() string {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v", APPID, APPSECRET)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("获取微信token失败", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("微信token读取失败", err)
		return ""
	}

	// token := make(map[string]interface{})
	token := token{}
	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Println("微信token解析json失败", err)
		return ""
	}

	return token.AccessToken
}

//获取关注人列表
func GetFlist(access_token string) []gjson.Result {
	url := "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" + access_token + "&next_openid="
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("获取关注列表失败", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取内容失败", err)
		return nil
	}
	flist := gjson.Get(string(body), "data.openid").Array()
	return flist
}

//发送模板消息代码
func templatepost(access_token string, reqdata string, fxurl string, templateid string, openid string) {
	url := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + access_token

	reqbody := "{\"touser\":\"" + openid + "\", \"template_id\":\"" + templateid + "\", \"url\":\"" + fxurl + "\", \"data\": " + reqdata + "}"

	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(string(reqbody)))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
func Test() {
	access_token := GetAccessToken()
	if access_token == "" {
		return
	}

	flist := GetFlist(access_token) //获取公众号关注人列表
	if flist == nil {
		return
	}
	a := "我是"
	b := "去"
	c := "往"

	reqdata := "{\"content\":{\"value\":\"jint:" + a + "\", \"color\":\"#0000CD\"}, \"note\":{\"value\":\"" + b + "\"}, \"translation\":{\"value\":\"" + c + "\"}}"
	fmt.Println(reqdata)
	for _, v := range flist {
		templatepost(access_token, reqdata, "", SentTemplateID, v.Str)
	}
}

func Send() {
	access_token := GetAccessToken()
	if access_token == "" {
		return
	}

	flist := GetFlist(access_token) //获取公众号关注人列表
	if flist == nil {
		return
	}
	weather, err := GetWeather()
	if err != nil {
		return
	}
	air, err := GetAir()
	if err != nil {
		return
	}
	talk, err := GetTalk()
	if err != nil {
		return
	}
	Birthday, err := GetBirthday()
	if err != nil {
		return
	}
	Love_day, err := GetLove()
	if err != nil {
		return
	}
	love_day := fmt.Sprintf("%d", Love_day)
	birthday := fmt.Sprintf("%d", Birthday)
	birthday_day := fmt.Sprintf("%d", Birthday-2)
	t := time.Now()
	date := fmt.Sprintf("%d年%d月%d日 %s", t.Year(), t.Month(), t.Day(), t.Weekday().String())
	name := []string{"date", "city", "weather", "temperature", "windspeed", "winddir", "air", "sunriseandset", "moonphase", "talk", "love_day", "birthday", "birthday_day"}
	value := []string{date, "湘潭", weather.TextDay, weather.TempMin + "°C-" + weather.TempMax + "°C", weather.WindSpeedDay, weather.WindDirDay, air.Category, weather.Sunrise + "," + weather.Sunset, weather.MoonPhase, talk, love_day, birthday, birthday_day}
	textdata := ""
	for i, l := range name {
		if i == 0 {
			textdata = textdata + `{"` + l + `":{"value":"` + value[i] + `","color":"` + GetRandomColor() + `"},`
		} else if i == len(name)-1 {
			textdata = textdata + `"` + l + `":{"value":"` + value[i] + `","color":"` + GetRandomColor() + `"}}`
		} else {
			textdata = textdata + `"` + l + `":{"value":"` + value[i] + `","color":"` + GetRandomColor() + `"},`
		}
	}
	fmt.Println(textdata)
	reqdata := `{"date":{"value":"` + date + `"},"city":{"value":"` + "湘潭" + `"},"weather":{"value":"` + weather.TextDay + `"},"temperature":{"value":"` + weather.TempMin + "°C-" + weather.TempMax + "°C" + `"},"windspeed":{"value":"` + weather.WindSpeedDay + `"},"winddir":{"value":"` + weather.WindDirDay + `"},"air":{"value":"` + air.Category + `"},"sunriseandset":{"value":"` + weather.Sunrise + "," + weather.Sunset + `"},"moonphase":{"value":"` + weather.MoonPhase + `"},"talk":{"value":"` + talk + `"},"love_day":{"value":"` + love_day + `"},"birthday":{"value":"` + birthday + `"},"birthday_day":{"value":"` + birthday_day + `"}}`
	fmt.Println(reqdata)
	for _, v := range flist {
		templatepost(access_token, textdata, "", SentTemplateID, v.Str)
	}
	return
}

func GetRandomColor() string {
	var letters = []rune("0123456789ABCDEF")
	color := "#"
	b := make([]rune, 6)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	color = color + string(b)
	return color
}

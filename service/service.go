package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/nosixtools/solarlunar"
)

type Weather struct {
	FxDate         string
	Sunrise        string
	Sunset         string
	Moonrise       string
	Moonset        string
	MoonPhase      string
	MoonPhaseIcon  string
	TempMax        string
	TempMin        string
	IconDay        string
	TextDay        string
	IconNight      string
	TextNight      string
	Wind360Day     string
	WindDirDay     string
	WindScaleDay   string
	WindSpeedDay   string
	Wind360Night   string
	WindDirNight   string
	WindScaleNight string
	WindSpeedNight string
	Humidity       string
	Precip         string
	Pressure       string
	Vis            string
	Cloud          string
	UvIndex        string
}

type Indices struct {
	Date     string
	Type     string
	Name     string
	Level    string
	Category string
	Text     string
}

type Air struct {
	FxDate   string
	Aqi      string
	Level    string
	Category string
	Primary  string
}

func GetWeather() (*Weather, error) {
	WeatherUrl := "https://devapi.qweather.com/v7/weather/3d?location=101250204&key=fe3ea7e00fd24669ae28ef3df5178215"
	resp, err := http.Get(WeatherUrl)
	if err != nil {
		return nil, err
	}
	WeatherUrl_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	weather := make(map[string]interface{})
	err = json.Unmarshal(WeatherUrl_body, &weather)
	if err != nil {
		return nil, err
	}
	collections := weather
	collectionInfo := &struct {
		Daily []Weather `json:"daily"`
	}{}
	collection, _ := json.Marshal(collections)
	err = json.Unmarshal(collection, &collectionInfo)
	if err != nil {
		return nil, err
	}
	return &collectionInfo.Daily[0], nil
}

func GetIndices() ([]Indices, error) {
	WeatherUrl := "https://devapi.qweather.com/v7/indices/1d?type=1,3,5,9,13,16&location=101250204&key=fe3ea7e00fd24669ae28ef3df5178215"
	resp, err := http.Get(WeatherUrl)
	if err != nil {
		return nil, err
	}
	WeatherUrl_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	weather := make(map[string]interface{})
	err = json.Unmarshal(WeatherUrl_body, &weather)
	if err != nil {
		return nil, err
	}
	collections := weather
	collectionInfo := &struct {
		Daily []Indices `json:"daily"`
	}{}
	collection, _ := json.Marshal(collections)
	err = json.Unmarshal(collection, &collectionInfo)
	if err != nil {
		return nil, err
	}
	return collectionInfo.Daily, nil
}

func GetAir() (*Air, error) {
	WeatherUrl := "https://devapi.qweather.com/v7/air/5d?location=101250204&key=fe3ea7e00fd24669ae28ef3df5178215"
	resp, err := http.Get(WeatherUrl)
	if err != nil {
		return nil, err
	}
	WeatherUrl_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	weather := make(map[string]interface{})
	err = json.Unmarshal(WeatherUrl_body, &weather)
	if err != nil {
		return nil, err
	}
	collections := weather
	collectionInfo := &struct {
		Daily []Air `json:"daily"`
	}{}
	collection, _ := json.Marshal(collections)
	err = json.Unmarshal(collection, &collectionInfo)
	if err != nil {
		return nil, err
	}
	return &collectionInfo.Daily[0], nil
}

func GetBirthday() (int, error) {
	timenow := time.Now()
	solarDate := timenow.Format("2006-01-02")
	year, _ := strconv.ParseInt(solarlunar.SolarToSimpleLuanr(solarDate)[:4], 10, 64)

	lunarDate := fmt.Sprintf("%d-01-15", year+1)
	solarBirthday := fmt.Sprintf("%s 00:00:00", solarlunar.LunarToSolar(lunarDate, false))
	time, _ := time.ParseInLocation("2006-01-02 15:04:05", solarBirthday, time.Local)
	day := int(time.Sub(timenow).Hours() / 24)
	return day, nil
}

func GetLove() (int, error) {
	timenow := time.Now()
	time, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-12-02 00:00:00", time.Local)
	day := int(timenow.Sub(time).Hours() / 24)
	return day, nil
}

func GetTalk() (string, error) {
	indices, err := GetIndices()
	if err != nil {
		return "", err
	}
	first := indices[3].Text[0:strings.Index(indices[3].Text, "。")]
	second := indices[1].Text[0:strings.Index(indices[1].Text, "。")]
	third := indices[4].Text[0:strings.Index(indices[4].Text, "。")]
	fourth := indices[5].Text
	return first + "，" + second + "。" + third + "。" + fourth, nil
}

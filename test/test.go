package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nosixtools/solarlunar"
)

func main() {
	timenow := time.Now()
	solarDate := timenow.Format("2006-01-02")
	fmt.Println(solarlunar.SolarToChineseLuanr(solarDate))
	year, _ := strconv.ParseInt(solarlunar.SolarToSimpleLuanr(solarDate)[:4], 10, 64)

	lunarDate := fmt.Sprintf("%d-01-15", year+1)
	fmt.Println(lunarDate)
	fmt.Println(solarlunar.LunarToSolar(lunarDate, false))
	solarBirthday := fmt.Sprintf("%s 00:00:00", solarlunar.LunarToSolar(lunarDate, false))
	time, _ := time.ParseInLocation("2006-01-02 15:04:05", solarBirthday, time.Local)
	day := int(time.Sub(timenow).Hours() / 24)
	fmt.Println(day)
}

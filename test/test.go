package main

import (
	"fmt"
	"weixin/service"
)

func main() {
	service.Send()
	a := `{"date":{"value":"2023年3月27日 Monday"},"city":{"value":"湘潭"},"weather":{"value":"小雨"},"temperature":{"value":"10°C-14°C"},"windspeed":{"value":"3"},"winddir":{"value":"北风"},"air":{"value":"良"},"sunriseandset":{"value":"06:25,18:45"},"moonphase":{"value":"峨眉月"},"talk":{"value":"天气较凉，较易发生感冒，请适当增加衣服，建议着厚外套加毛衣等服装。皮肤易缺水，用润唇膏后再抹口红，用保湿型霜类化妆品。属弱紫外辐射天气，长期在户外，建议涂擦SPF在8-12之间的防晒护肤品。"},"love_day":{"value":"115"},"birthday":{"value":"333"},"birthday_day":{"value":"331"}}`
	b := `{"date":{"value":"2023年3月27日 Monday"},"city":{"value":"湘潭"},"weather":{"value":"小雨"},"temperature":{"value":"10°C-14°C"},"windspeed":{"value":"3"},"winddir":{"value":"北风"},"air":{"value":"良"},"sunriseandset":{"value":"06:25,18:45"},"moonphase":{"value":"峨眉月"},"talk":{"value":"天气较凉，较易发生感冒，请适当增加衣服，建议着厚外套加毛衣等服装。皮肤易缺水，用润唇膏后再抹口红，用保湿型霜类化妆品。属弱紫外辐射天气，长期在户外，建议涂擦SPF在8-12之间的防晒护肤品。"},"love_day":{"value":"115"},"birthday":{"value":"333"},"birthday_day":{"value":"331"}}`
	if a == b {
		fmt.Println("sss")
	}
}

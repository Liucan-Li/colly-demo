package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	// 创建Collector实例
	c := colly.NewCollector()

	// 设置请求头
	headers := map[string][]string{
		"accept":             []string{"application/json, text/plain, */*"},
		"accept-language":    []string{"zh-CN,zh;q=0.9,en;q=0.8,und;q=0.7"},
		"content-type":       []string{"application/json"},
		"cookie":             []string{"tracking-preferences={%22version%22:1%2C%22destinations%22:{%22Actions%20Amplitude%22:true%2C%22AdWords%22:true%2C%22Google%20AdWords%20New%22:true%2C%22Google%20Enhanced%27Conversions%22:true%2C%22Google%20Tag%20Manager%22:true%2C%22Impact%20Partnership%20Cloud%22:true}%2C%22custom%22:{%22advertising%22:true%2C%22functional%22:true%2C%22marketingAndAnalytics%22:true}}; ajs_anonymous_id=8d02d26c-7b9e-4585-853b-182e643d2ecc; _gcl_au=1.1.2113232548.1765094966; _ga=GA1.1.440162287.1765094967; TCG_VisitorKey=183812f7-e6dd-4f67-bc88-224effdec23d; setting=CD=CN&M=1; SellerProximity=ZipCode=&MaxSellerDistance=1000&IsActive=false; SearchSortSettings=M=1&ProductSortOption=BestMatch&ProductSortDesc=False&PriceSortOption=Shipping&ProductResultDisplay=grid; analytics_session_id=1765893720049; _ga_JEQYTNS2WQ=GS2.1.s1765893734$o4$g0$t1765893734$j60$l0$h0; _ga_KK8XBGNYRB=GS2.1.s1765893734$o4$g0$t1765893734$j60$l0$h0; _ga_0T2XGBC5QN=GS2.1.s1765893734$o4$g0$t1765893734$j60$l0$h0; tcg-segment-session=1765893708684%257C1765893834560; analytics_session_id.last_access=1765893835356; _ga_VS9BE2Z3GY=GS2.1.s1765893721$o4$g1$t1765893835$j60$l0$h727653683"},
		"origin":             []string{"https://www.tcgplayer.com"},
		"priority":           []string{"u=1, i"},
		"referer":            []string{"https://www.tcgplayer.com/"},
		"sec-ch-ua":          []string{`"Google Chrome";v="143", "Chromium";v="143", "Not A(Brand";v="24"`},
		"sec-ch-ua-mobile":   []string{"?0"},
		"sec-ch-ua-platform": []string{`"Linux"`},
		"sec-fetch-dest":     []string{"empty"},
		"sec-fetch-mode":     []string{"cors"},
		"sec-fetch-site":     []string{"same-site"},
		"user-agent":         []string{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36"},
	}

	// 请求体数据
	requestBody := `{"algorithm":"sales_dismax","from":24,"size":24,"filters":{"term":{},"range":{},"match":{}},"listingSearch":{"context":{"cart":{}},"filters":{"term":{"sellerStatus":"Live","channelId":0},"range":{"quantity":{"gte":1}},"exclude":{"channelExclusion":0}}},"context":{"cart":{},"shippingCountry":"CN","userProfile":{}},"settings":{"useFuzzySearch":true,"didYouMean":{}},"sort":{}}`

	// 设置响应处理回调
	c.OnResponse(func(r *colly.Response) {
		log.Printf("响应状态码: %d\n", r.StatusCode)
		log.Printf("响应体: %s\n", string(r.Body))
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("请求错误: %v\n", err)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Printf("正在访问: %s\n", r.URL.String())
	})

	// 发送POST请求
	err := c.Request(
		"POST",
		"https://mp-search-api.tcgplayer.com/v1/search/request?q=&isList=false&mpfev=4616",
		strings.NewReader(requestBody),
		nil,
		http.Header(headers),
	)

	if err != nil {
		log.Fatal(err)
	}

	// 等待请求完成
	c.Wait()
}

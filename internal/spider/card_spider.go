package spider

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Liucan-Li/colly-demo/internal/dao"
	"github.com/Liucan-Li/colly-demo/internal/dao/model"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type CardSpider struct{}

var Spider = &CardSpider{}

type CardInfo struct{}
type Result struct {
	Results []map[string]interface{}
}

type SpiderResponse struct {
	Errors  []string
	Results []Result
}

func (s *CardSpider) DoScrawling(page int) {
	// 创建Collector实例
	c := colly.NewCollector()
	// 设置随机UserAgent
	extensions.RandomUserAgent(c)
	// 设置请求头
	headers := map[string][]string{
		"accept":             {"application/json, text/plain, */*"},
		"accept-language":    {"zh-CN,zh;q=0.9,en;q=0.8,und;q=0.7"},
		"content-type":       {"application/json"},
		"cookie":             {`TCG_VisitorKey=04c2b2e7-7277-49c2-a57f-5c9202d1b3fd; tracking-preferences={%22version%22:1%2C%22destinations%22:{%22Actions%20Amplitude%22:true%2C%22AdWords%22:true%2C%22Google%20AdWords%20New%22:true%2C%22Google%20Enhanced%20Conversions%22:true%2C%22Google%20Tag%20Manager%22:true%2C%22Impact%20Partnership%20Cloud%22:true}%2C%22custom%22:{%22advertising%22:true%2C%22functional%22:true%2C%22marketingAndAnalytics%22:true}}; SellerProximity=ZipCode=&MaxSellerDistance=1000&IsActive=false; SearchSortSettings=M=1&ProductSortOption=BestMatch&ProductSortDesc=False&PriceSortOption=Shipping&ProductResultDisplay=grid; ajs_anonymous_id=789097fe-4e9b-4739-b384-a460e3e0d195; _gcl_au=1.1.1780692817.1766765206; _ga=GA1.1.135138769.1766765206; analytics_session_id=1766798327206; setting=CD=HK&M=1; SearchCriteria=M=1&WantVerifiedSellers=False&WantDirect=False&WantSellersInCart=False&WantWPNSellers=False; tcg-segment-session=1766798327208%257C1766798331008; analytics_session_id.last_access=1766798331081; _ga_VS9BE2Z3GY=GS2.1.s1766798328$o2$g1$t1766798331$j57$l0$h208256392`},
		"origin":             {"https://www.tcgplayer.com"},
		"priority":           {"u=1, i"},
		"referer":            {"https://www.tcgplayer.com/"},
		"sec-ch-ua":          {`"Google Chrome";v="143", "Chromium";v="143", "Not A(Brand";v="24"`},
		"sec-ch-ua-mobile":   {"?0"},
		"sec-ch-ua-platform": {`"Linux"`},
		"sec-fetch-dest":     {"empty"},
		"sec-fetch-mode":     {"cors"},
		"sec-fetch-site":     {"same-site"},
		// "user-agent":         {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36"},
	}

	size := 24
	from := page * size

	// 请求体数据
	requestBody := fmt.Sprintf(`{"algorithm":"sales_dismax","from":%d,"size":%d,"filters":{"term":{},"range":{},"match":{}},"listingSearch":{"context":{"cart":{}},"filters":{"term":{"sellerStatus":"Live","channelId":0},"range":{"quantity":{"gte":1}},"exclude":{"channelExclusion":0}}},"context":{"cart":{},"shippingCountry":"CN","userProfile":{}},"settings":{"useFuzzySearch":true,"didYouMean":{}},"sort":{}}`, from, size)

	// 设置响应处理回调
	c.OnResponse(func(r *colly.Response) {
		var res SpiderResponse
		if err := json.Unmarshal(r.Body, &res); err != nil {
			log.Printf("解析错误：%v", err)
			return
		}

		if res.Results != nil {
			for _, result := range res.Results {
				if len(result.Results) == 0 {
					continue
				}
				for _, contentInfo := range result.Results {
					contentStr, _ := json.Marshal(contentInfo)
					v := &model.TcgCardOrigin{
						OriginContent: string(contentStr),
					}

					dao.TcgCard.Create(v)
				}
			}
		}

		log.Printf("请求成功: %d", page)

	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("请求错误: %v\n; page: %d", err, page)
		s.DoScrawling(page)
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
		log.Printf("抓取错误: %v\n; page: %d", err, page)
	}

	// 等待请求完成
	c.Wait()
}

func (s *CardSpider) DoBatchScrawling() {
	for i := 0; i < 10; i++ {
		s.DoScrawling(i)
	}
}

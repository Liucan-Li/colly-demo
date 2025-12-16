// package main

// import (
// 	"context"
// 	"log"
// 	"strings"

// 	"github.com/chromedp/chromedp"
// 	"github.com/gocolly/colly"
// )

// func main() {

// 	ctx, cancel := chromedp.NewContext(context.Background())
// 	defer cancel()

// 	var finalHTML string // 用于保存JS执行后的完整HTML

// 	err := chromedp.Run(ctx,
// 		chromedp.Navigate(`https://www.tcgplayer.com/search/all/product?view=grid`),
// 		chromedp.WaitReady(`.product-card__product product-card__product-variant-a`), // 等待动态内容加载完成的关键选择器
// 		chromedp.OuterHTML("html", &finalHTML),                                       // 获取整个文档的HTML
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	c := colly.NewCollector()
// 	c.OnHTML(".product-card__product product-card__product-variant-a", func(e *colly.HTMLElement) {
// 		// 现在可以像解析静态页面一样提取数据了
// 		title := e.ChildText("h3")
// 		log.Printf("标题: %s\n", title)
// 	})

// 	// 技巧：将HTML字符串"导入"Colly进行处理
// 	err = c.ParseReader(strings.NewReader(finalHTML))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// // Instantiate default collector
// 	// c := colly.NewCollector(
// 	// // Visit only domains: hackerspaces.org, wiki.hackerspaces.org
// 	// // colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
// 	// )

// 	// extensions.RandomUserAgent(c)

// 	// // On every a element which has href attribute call callback
// 	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 	// 	link := e.Attr("href")
// 	// 	// Print link
// 	// 	fmt.Printf("Link found: %q -> %s\n", e.Text, link)
// 	// 	// Visit link found on page
// 	// 	// Only those links are visited which are in AllowedDomains
// 	// 	c.Visit(e.Request.AbsoluteURL(link))
// 	// })

// 	// // Before making a request print "Visiting ..."
// 	// c.OnRequest(func(r *colly.Request) {
// 	// 	fmt.Println("Visiting", r.URL.String())
// 	// })

// 	// c.OnResponse(func(r *colly.Response) {
// 	// 	fmt.Println("%v: \n", string(r.Body))
// 	// })
// 	// // Set error handler
// 	// c.OnError(func(r *colly.Response, err error) {
// 	// 	fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
// 	// })

// 	// // Start scraping on https://hackerspaces.org
// 	// c.Visit("https://www.tcgplayer.com/search/all/product?view=grid")
// 	// // c.Visit("https://en.wikipedia.org/")

// }

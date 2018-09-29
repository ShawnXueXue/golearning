package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"net/url"
	"time"
)

func main() {
	//simpleColly()
	collyAndGoquery()
}

func simpleColly() {
	c := colly.NewCollector()
	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		element.Request.Visit(element.Attr("href"))
	})
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL)
	})
	c.OnError(func(response *colly.Response, e error) {
		fmt.Println("Error:", e)
	})
	c.Visit("http://go-colly.org/")
}

func collyAndGoquery() {
	urlStr := "http://metalsucks.net"
	u, err := url.Parse(urlStr)
	if err != nil {
		log.Fatal(err)
	}
	c := colly.NewCollector()
	c.SetRequestTimeout(100 * time.Second)
	// 指定agent信息
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36"
	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("Host", u.Host)
		request.Headers.Set("Connection", "keep-alive")
		request.Headers.Set("Accept", "*/*")
		request.Headers.Set("Origin", u.Host)
		request.Headers.Set("Referer", urlStr)
		request.Headers.Set("Accept-Encoding", "gzip, deflate")
		request.Headers.Set("Accept-language", "zh-CN, zh;q=0.9")
	})
	c.OnResponse(func(response *colly.Response) {
		// colly读取的内容传给goquery
		htmlDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(response.Body))
		if err != nil {
			log.Fatal(err)
		}
		// 找到抓取项
		htmlDoc.Find(".sidebar-reviews article .content-block").Each(func(i int, selection *goquery.Selection) {
			band := selection.Find("a").Text()
			title := selection.Find("i").Text()
			fmt.Printf("Review %d: %s - %s\n", i, band, title)
		})
	})
	c.OnError(func(response *colly.Response, e error) {
		err = e
	})
	err = c.Visit(urlStr)
}

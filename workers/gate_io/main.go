package main

import (
	"fmt"
	"time"

	// "github.com/apex/log"
	"runtime/debug"

	"log"

	"github.com/gocolly/colly/v2"
)

type Item struct {
	StoryURL string
	Source   string
	ReNewAt  string
	Comments string
	Title    string
}

func main() {
	startUrl := "https://www.gate.io/cn/articlelist/ann/0"

	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36"),
	)

	c.OnHTML(".leftlatnews", func(e *colly.HTMLElement) {
		item := Item{}
		item.Title = e.ChildText(".entry > a > h3")
		href := e.ChildAttr(".entry > a", "href")
		item.StoryURL = "https://www.gate.io" + href
		item.ReNewAt = time.Now().Format("2006-01-02 15:04:05")
		item.Source = "芝麻开门"
		fmt.Println(item)
		// if err := crawlab.SaveItem(item); err != nil {
		//     log.Errorf("save item error: " + err.Error())
		//     debug.PrintStack()
		//     return
		// }
	})

	c.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting %s", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	if err := c.Visit(startUrl); err != nil {
		log.Fatalln("visit error: " + err.Error())
		debug.PrintStack()
		panic(fmt.Sprintf("Unable to visit %s", startUrl))
	}

	c.Wait()
}

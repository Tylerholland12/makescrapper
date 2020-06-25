package main

import (
	"log"
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })
	

	c.OnHTML("tr td:nth-of-type(2)", func(e *colly.HTMLElement) {
		fmt.Println("The name of the Cryptocurrency is:", e.Text)
	})

	c.OnHTML("tr td:nth-of-type(3)", func(e *colly.HTMLElement) {
		fmt.Println("The price of the Cryptocurrency is:", e.Text)
	})
	
	c.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println(e.Text)
	})
	
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	
	c.Visit("https://finance.yahoo.com/cryptocurrencies")
}
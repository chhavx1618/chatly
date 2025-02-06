package main 

import (
	"fmt"
	"log"
	"os"
	"github.com/gocolly/colly"
)


type Product struct {
	url, img, name, price string
}


func main () {
	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com"),
	)
	//trial - scraping course

	var products []product

	product.url = e.ChildAttr("a", "href")

	c.OnHTML("a", func(e *colly.HTMLElement) {

		product := product{}
		fmt.Print("%v", e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Print("visiting ", r.URL)
	})

	c.OnScraped(func(r *colly.Response) {
        fmt.Println(r.Request.URL, " scraped!")
    })

	c.Visit("amazon.com")
}
package main

import (
	"fmt"
	// "net/http"
	// "log"
	"github.com/gocolly/colly"
)

type Apartment struct {
	Location string
	Sqft string
	Price string
	NumBedrooms string
}

// func scrapeHandler() {
// 	// Instantiate default collector
// 	c := colly.NewCollector()

// 	// On every a element which has href attribute call callback
// 	c.OnHTML(".list-card-price", func(e *colly.HTMLElement) {
// 		fmt.Printf("Price per month: %q\n", e.Text)
// 	})

// 	c.OnHTML(".list-card-details li:first-child", func(e *colly.HTMLElement) {
// 		fmt.Printf("Number of bedrooms: %q\n", e.Text)
// 	})

// 	c.OnHTML(".list-card-addr", func(e *colly.HTMLElement) {
// 		link := e.Attr("href")
// 		fmt.Printf("Location: %q\nLink: %s\n", e.Text, link)
// 	})

// 	c.OnHTML(".list-card-details li:nth-child(3)", func(e *colly.HTMLElement) {
// 		fmt.Printf("Sqft: %q\n", e.Text)
// 	})

// 	// Before making a request print "Visiting ..."
// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL.String())
// 	})

// 	// Start scraping on https://www.zillow.com/homes/for_rent/
// 	c.Visit("https://www.zillow.com/homes/for_rent/")
// }

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	// c.OnHTML(".list-card-price", func(e *colly.HTMLElement) {
	// 	fmt.Printf("Price per month: %q\n", e.Text)
	// })
	c.OnHTML(".list-card-info", func(e *colly.HTMLElement) {
		fmt.Printf("Full info card: %q\n\n", e.Text)
	})

	// c.OnHTML(".list-card-details li:first-child", func(e *colly.HTMLElement) {
	// 	fmt.Printf("Number of bedrooms: %q\n", e.Text)
	// })

	// c.OnHTML(".list-card-addr", func(e *colly.HTMLElement) {
	// 	link := e.Attr("href")
	// 	fmt.Printf("Location: %q\nLink: %s\n", e.Text, link)
	// })

	// c.OnHTML(".list-card-details li:nth-child(3)", func(e *colly.HTMLElement) {
	// 	fmt.Printf("Sqft: %q\n", e.Text)
	// })

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	
	c.Visit("https://www.zillow.com/homes/San-Francisco,-CA_rb/?searchQueryState=%7B%22pagination%22%3A%7B%7D%2C%22usersSearchTerm%22%3A%22San%20Francisco%2C%20CA%22%2C%22mapBounds%22%3A%7B%22west%22%3A-122.54679717016602%2C%22east%22%3A-122.31986082983398%2C%22south%22%3A37.68527794262908%2C%22north%22%3A37.8651955033787%7D%2C%22regionSelection%22%3A%5B%7B%22regionId%22%3A20330%2C%22regionType%22%3A6%7D%5D%2C%22filterState%22%3A%7B%22pmf%22%3A%7B%22value%22%3Afalse%7D%2C%22fore%22%3A%7B%22value%22%3Afalse%7D%2C%22auc%22%3A%7B%22value%22%3Afalse%7D%2C%22nc%22%3A%7B%22value%22%3Afalse%7D%2C%22fr%22%3A%7B%22value%22%3Atrue%7D%2C%22fsbo%22%3A%7B%22value%22%3Afalse%7D%2C%22cmsn%22%3A%7B%22value%22%3Afalse%7D%2C%22pf%22%3A%7B%22value%22%3Afalse%7D%2C%22fsba%22%3A%7B%22value%22%3Afalse%7D%7D%2C%22isListVisible%22%3Atrue%2C%22isMapVisible%22%3Atrue%2C%22mapZoom%22%3A12%7D")

	// Start scraping on https://www.zillow.com/homes/for_rent/
	// c.Visit("https://www.zillow.com/homes/for_rent/")
	// host := "0.0.0.0:8888"
	// http.HandleFunc("/", scrapeHandler)

	// fmt.Println("Starting server: http://" + host)
	// err := http.ListenAndServe(host, nil)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
}

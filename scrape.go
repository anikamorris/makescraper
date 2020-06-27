package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Apartment struct {
	Location string
	Sqft string
	Price string
	NumBedrooms string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML(".list-card-price", func(e *colly.HTMLElement) {
		fmt.Printf("Price per month: %q\n", e.Text)
	})

	c.OnHTML(".list-card-details li:first-child", func(e *colly.HTMLElement) {
		fmt.Printf("Number of bedrooms: %q\n", e.Text)
	})

	c.OnHTML(".list-card-addr", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Location: %q\nLink: %s\n", e.Text, link)
	})

	c.OnHTML(".list-card-details li:nth-child(3)", func(e *colly.HTMLElement) {
		fmt.Printf("Sqft: %q\n", e.Text)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://www.zillow.com/homes/for_rent/
	c.Visit("https://www.zillow.com/homes/for_rent/")
}

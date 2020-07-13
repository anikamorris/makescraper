package main

import (
	"fmt"
	// "net/http"
	"log"
	// "strings"
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/gocolly/colly"
)

/* Apartment struct to keep track of address, square footage, 
   price per month, and number of bedrooms */
type Apartment struct {
	Location string
	Sqft string
	Price string
	NumBedrooms string
}

/* struct to store the full listing before being processed */
type FullListing struct {
	Listing string
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

/* Writes listing to inputted file */
func AppendListingToFile(filename string, e FullListing) {
	listingJSON, err := json.Marshal(e)
	if err != nil {
		log.Fatalf("failed to encode listing as json")
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    defer file.Close()

	jsonString := string(listingJSON)
    _, err = file.WriteString(jsonString + ",\n")
    if err != nil {
        log.Fatalf("failed writing to file: %s", err)
    }
}

func main() {

	c := colly.NewCollector()
	// Keep track of all listings that we've found so far
	var listings []FullListing
	// On every a element which has href attribute call callback
	// c.OnHTML(".list-card-price", func(e *colly.HTMLElement) {
	// 	fmt.Printf("Price per month: %q\n", e.Text)
	// })
	c.OnHTML(".list-card-info", func(e *colly.HTMLElement) {
		listing := FullListing{Listing: e.Text}
		listings = append(listings, listing)
		lJson, _ := json.Marshal(listing)
		_ = ioutil.WriteFile("output.json", lJson, 0644)
		
		for _, listing := range listings {
			AppendListingToFile("output.json", listing)
		}
		
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

	// Handle errors
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
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

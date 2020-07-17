package main

import (
	"fmt"
	// "net/http"
	"log"
	"strings"
	"strconv"
	"os"
	// "io/ioutil"
	"encoding/json"
	"github.com/gocolly/colly"
	"regexp"
)

/* Apartment struct to keep track of address, square footage, 
   price per month, and number of bedrooms */
type Apartment struct {
	Location string
	Price string
	Details string
	URL string
}

/* struct to store the full listing before being processed */
type FullListing struct {
	Listing string
}

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

func homeView() []Apartment {
	c := colly.NewCollector()
	// Keep track of all listings that we've found so far
	var listings []Apartment
	// On every a element which has href attribute call callback
	// c.OnHTML(".list-card-price", func(e *colly.HTMLElement) {
	// 	fmt.Printf("Price per month: %q\n", e.Text)
	// })
	c.OnHTML(".list-card-link", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %s -> %s \n", e.Text, link)
		apartment := detailView(e.Text, link)
		listings = append(listings, apartment)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Handle errors
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.Visit("https://www.zillow.com/homes/San-Francisco,-CA_rb/?searchQueryState=%7B%22pagination%22%3A%7B%7D%2C%22usersSearchTerm%22%3A%22San%20Francisco%2C%20CA%22%2C%22mapBounds%22%3A%7B%22west%22%3A-122.54679717016602%2C%22east%22%3A-122.31986082983398%2C%22south%22%3A37.68527794262908%2C%22north%22%3A37.8651955033787%7D%2C%22regionSelection%22%3A%5B%7B%22regionId%22%3A20330%2C%22regionType%22%3A6%7D%5D%2C%22filterState%22%3A%7B%22pmf%22%3A%7B%22value%22%3Afalse%7D%2C%22fore%22%3A%7B%22value%22%3Afalse%7D%2C%22auc%22%3A%7B%22value%22%3Afalse%7D%2C%22nc%22%3A%7B%22value%22%3Afalse%7D%2C%22fr%22%3A%7B%22value%22%3Atrue%7D%2C%22fsbo%22%3A%7B%22value%22%3Afalse%7D%2C%22cmsn%22%3A%7B%22value%22%3Afalse%7D%2C%22pf%22%3A%7B%22value%22%3Afalse%7D%2C%22fsba%22%3A%7B%22value%22%3Afalse%7D%7D%2C%22isListVisible%22%3Atrue%2C%22isMapVisible%22%3Atrue%2C%22mapZoom%22%3A12%7D")
	return listings
}

func detailView(location string, link string) Apartment {
	var apartment Apartment
	apartment.Location = location
	apartment.URL = link

	c := colly.NewCollector()

	c.OnHTML(".ds-summary-row", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildText(".ds-value"))
		fmt.Println(e.ChildText(".ds-bed-bath-living-area"))
		apartment.Details = e.ChildText(".ds-bed-bath-living-area")
		apartment.Price = e.ChildText(".ds-value")
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("-------------------")
		fmt.Println("Visiting: " + location)
	})

	c.Visit(link)
	return apartment
}

func filterApartmentsByPrice(apartments []Apartment, price int) []Apartment {
	fmt.Println(price)
	var apartmentsInPriceRange []Apartment
	for _, apartment := range apartments {
		if apartment.Price != "" {
			aPrice := strings.Split(apartment.Price, "$")
			fmt.Printf("apartment price: %s \n", aPrice[1])
			reg, err := regexp.Compile("[^a-zA-Z0-9]+")
			if err != nil {
				log.Fatal(err)
			}
			processedPrice := reg.ReplaceAllString(aPrice[1], "")
			intPrice, _ := strconv.Atoi(processedPrice)
			// fmt.Printf("Int price: %s \n", priceNoComma)
			if intPrice <= price {
				fmt.Println(intPrice)
				apartmentsInPriceRange = append(apartmentsInPriceRange, apartment)
			}
		}
	}
	return apartmentsInPriceRange
}

func main() {
	apartments := homeView()
	apartmentsInPriceRange := filterApartmentsByPrice(apartments, 4000)
	fmt.Println(apartmentsInPriceRange)
	// filterApartmentsByPrice(apartments, 4000)
}

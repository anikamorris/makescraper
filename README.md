# 🕷 makeHousing

[![Go Report Card](https://goreportcard.com/badge/github.com/anikamorris/makescraper)](https://goreportcard.com/report/github.com/anikamorris/makescraper)

_Create your very own web scraper and crawler using Go and [Colly](https://go-colly.org)!_

### 📚 Table of Contents

1. [Project Structure](#project-structure)
2. [Usage](#usage)
3. [Some Background](#some-background)
4. [Resources](#resources)

## Project Structure

```bash
📂 makescraper
├── output.json
├── README.md
└── scrape.go
```

## Usage
Clone and cd into directory. Run `go build`, then `./makescraper`. If you would like to specify a maximum price, add and set a price flag when running ./makescraper. For example, if you wanted your maximum price to be 4000, it would look like `./makescraper -price=4000`

## Some Background
https://docs.google.com/presentation/d/14dSsrSytvioamio81eZ_juF1LEh_9xBuFMbOKNL5bWk/edit?usp=sharing

## Resources
#### Scraping

- [**Colly** - Docs](http://go-colly.org/docs/): Check out the sidebar for 20+ examples!
- [**Ali Shalabi** - Syntax-Helper](https://github.com/alishalabi/syntax-helper): Command line interface to help generate proper code syntax, pulled from the Golang documentation.

#### Serializing & Saving

- [JSON to Struct](https://mholt.github.io/json-to-go/): Paste any JSON data and convert it into a Go structure that will support storing that data.
- [GoByExample - JSON](https://gobyexample.com/json): Covers Go's built-in support for JSON encoding and decoding to and from built-in and custom data types (structs).
- [GoByExample - Writing Files](https://gobyexample.com/writing-files): Covers creating new files and writing to them.

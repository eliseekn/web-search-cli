package search

import (
	"fmt"
	"web-search-cli/scraper"
)

func Init(engine string, query string, pages int) {
	if engine == "ddgo" {
		links := scraper.DDGoScrapper(query)

		for _, link := range links {
			fmt.Println(link)
		}

		return
	}

	for i := 1; i <= pages; i++ {
		links := scraper.GoogleScraper(query, i*10)

		for _, link := range links {
			fmt.Println(link)
		}
	}
}

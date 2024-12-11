package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

func GoogleScraper(query string, start int) []string {
	var links []string
	c := Init()

	c.OnHTML("a", func(e *colly.HTMLElement) {
		text := e.ChildText("h3")
		link := e.Attr("href")

		if text != "" {
			formattedLink := formatLink(link, "/url?q=", "&sa=U&ved=")
			links = append(links, fmt.Sprintf("[-] %v [%v]", text, formattedLink))
		}
	})

	c.Visit("https://google.com/search?hl=en&q=" + formatQuery(query) + "&start=" + fmt.Sprintf("%v", start))

	return links
}

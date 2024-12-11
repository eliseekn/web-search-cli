package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

func DDGoScrapper(query string) []string {
	var links []string
	c := Init()

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		text := e.ChildText("a")
		link := e.ChildAttr("a", "href")

		if text != "" {
			formattedLink := formatLink(link, "//duckduckgo.com/l/?uddg=", "&rut=")
			links = append(links, fmt.Sprintf("[-] %v [%v]", text, formattedLink))
		}
	})

	c.Visit("https://html.duckduckgo.com/html/?q=" + formatQuery(query))

	return links
}

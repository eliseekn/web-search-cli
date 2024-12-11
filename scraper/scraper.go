package scraper

import (
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

func Init() *colly.Collector {
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		panic(err)
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")
	})

	return c
}

func formatLink(link string, prefix string, suffix string) string {
	link, err := url.QueryUnescape(link)
	if err != nil {
		panic(err)
	}

	link = strings.ReplaceAll(link, prefix, "")

	index := strings.Index(link, suffix)
	if index != -1 {
		link = link[:index]
	}

	return link
}

func formatQuery(query string) string {
	var result string

	if strings.Contains(query, ":") {
		return url.QueryEscape(query)
	}

	for _, q := range strings.Split(query, " ") {
		result = result + q + "+"
	}

	return result
}

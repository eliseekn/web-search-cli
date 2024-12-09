package scraper

import (
	"net/url"
	"strings"
)

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

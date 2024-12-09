package main

import (
	"flag"
	"web-search-cli/search"
)

func main() {
	query := flag.String("q", "", "search query")
	pages := flag.Int("p", 1, "nombre of pages results (only for google engine)")
	engine := flag.String("e", "ddgo", "search engine to use")

	flag.Parse()

	search.Init(*engine, *query, *pages)
}

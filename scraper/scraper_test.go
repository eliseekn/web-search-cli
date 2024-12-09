package scraper_test

import (
	"testing"
	"web-search-cli/scraper"
)

func Test_Can_Get_Correct_DuckduckGo_Results_Length(t *testing.T) {
	links := scraper.DDGoScrapper("learn golang")

	if len(links) != 10 {
		t.FailNow()
	}
}

func Test_Can_Get_Correct_Google_Results_Length(t *testing.T) {
	links := scraper.GoogleScraper("learn golang", 0)

	if len(links) != 10 {
		t.FailNow()
	}
}

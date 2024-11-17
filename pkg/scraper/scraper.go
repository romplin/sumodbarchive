package scraper

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

func NewScraper() (*Scraper, error) {
	return &Scraper{
		baseURL: "https://sumodb.sumogames.de",
	}, nil
}

func (s *Scraper) Run() error {
	c := colly.NewCollector(
		colly.AllowedDomains("sumodb.sumogames.de"),
		colly.UserAgent("Mozilla/5.0 (compatible; SumoBot/1.0; +http://example.com)"),
	)

	// Set up rate limiting
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*sumodb.sumogames.de*",
		Parallelism: 1,
		Delay:       2 * time.Second,
	})

	// Handle errors
	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Request URL: %v failed with response: %v\nError: %v\n", r.Request.URL, r, err)
	})

	// Log visits
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Parse wrestler data
	c.OnHTML("table.tk_table", func(e *colly.HTMLElement) {
		wrestler := SumoWrestler{}

		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			wrestler.Name = row.ChildText("td:nth-child(2)")
			wrestler.Rank = row.ChildText("td:nth-child(3)")

			if wrestler.Name != "" {
				fmt.Printf("Found wrestler: %+v\n", wrestler)
			}
		})
	})

	return c.Visit(s.baseURL + "/Default.aspx")
}

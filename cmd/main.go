package main

import (
	"log"

	"sumodbarchive/pkg/scraper"
)

func main() {
	scraper, err := scraper.NewScraper()
	if err != nil {
		log.Fatalf("Failed to initialize scraper: %v", err)
	}

	if err := scraper.Run(); err != nil {
		log.Fatalf("Error during scraping: %v", err)
	}
}

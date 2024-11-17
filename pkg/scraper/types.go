package scraper

type SumoWrestler struct {
	Name       string
	Rank       string
	Wins       int
	Losses     int
	Tournament string
}

type Scraper struct {
	baseURL string
}

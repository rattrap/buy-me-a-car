package crawlers

import (
	"net/url"

	"github.com/rattrap/buy-me-a-car/model"
)

type genericCrawler interface {
	Apply(*url.URL) []model.Car
}

// Crawl is a
type Crawl struct {
	Crawler genericCrawler
}

// Crawl is used for parsing
func (o *Crawl) Crawl(url *url.URL) []model.Car {
	return o.Crawler.Apply(url)
}

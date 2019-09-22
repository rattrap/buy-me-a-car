package crawlers

import (
	"fmt"
	"net/url"

	"github.com/gocolly/colly"
	"github.com/rattrap/buy-me-a-car/model"
)

// Olx is the olx.ro crawler implementation
type Olx struct{}

// Apply is the olx.ro crawler implementation
func (Olx) Apply(url *url.URL) []model.Car {
	cars := []model.Car{}
	c := colly.NewCollector(
		colly.AllowedDomains("www.olx.ro"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	detailCollector := c.Clone()

	c.OnHTML("td.offer", func(e *colly.HTMLElement) {
		link := e.ChildAttr("h3 a", "href")
		detailCollector.Visit(link)
	})

	detailCollector.OnHTML("body", func(e *colly.HTMLElement) {
		car := model.Car{
			URL:     e.Request.URL.String(),
			Title:   e.ChildText(".offer-titlebox h1"),
			Picture: e.ChildAttr("img.bigImage", "src"),
		}

		e.ForEach("table.details > tbody > tr > td", func(_ int, el *colly.HTMLElement) {
			switch el.ChildText("th") {
			case "An de fabricatie":
				car.Year = el.ChildText("td")
			case "Rulaj":
				car.Km = el.ChildText("td")
			}
		})
		cars = append(cars, car)
	})

	c.Visit(url.String())
	return cars
}

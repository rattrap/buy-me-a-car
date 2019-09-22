package crawlers

import (
	"fmt"
	"net/url"

	"github.com/gocolly/colly"
	"github.com/rattrap/buy-me-a-car/model"
)

// Carzz is the carzz.ro crawler implementation
type Carzz struct{}

// Apply is the carzz.ro crawler implementation
func (Carzz) Apply(url *url.URL) []model.Car {
	cars := []model.Car{}
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	detailCollector := c.Clone()

	c.OnHTML("#list_cart_holder a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		detailCollector.Visit(link)
	})

	detailCollector.OnHTML("body", func(e *colly.HTMLElement) {
		car := model.Car{
			URL:     e.Request.URL.String(),
			Title:   e.ChildText("h1"),
			Picture: e.ChildAttr("img[itemprop=representativeOfPage]", "src"),
		}
		e.ForEach("#extra-fields > div", func(_ int, el *colly.HTMLElement) {
			switch el.ChildText("span") {
			case "An fabricație":
				car.Year = el.ChildText("h5")
			case "Rulaj":
				car.Km = el.ChildText("h6")
			}
		})
		cars = append(cars, car)
	})

	c.Visit(url.String())
	return cars
}

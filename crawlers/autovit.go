package crawlers

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
	"github.com/rattrap/buy-me-a-car/model"
)

// Autovit is the autovit.ro crawler implementation
type Autovit struct{}

// Apply is the autovit.ro crawler implementation
func (Autovit) Apply(url *url.URL) []model.Car {
	cars := []model.Car{}
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML(".offers > article", func(e *colly.HTMLElement) {
		car := model.Car{
			URL:     e.Attr("data-href"),
			Title:   e.ChildText(".offer-title"),
			Picture: strings.Split(e.ChildAttr("a.ds-photo img", "data-src"), ";")[0],
		}

		e.ForEach(".ds-params-block > li", func(_ int, el *colly.HTMLElement) {
			switch el.Attr("data-code") {
			case "year":
				car.Year = el.ChildText("span")
			case "mileage":
				car.Km = el.ChildText("span")
			}
		})
		cars = append(cars, car)
	})

	c.Visit(url.String())
	return cars
}

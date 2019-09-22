package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/url"
	"os"
	s "strings"

	"github.com/rattrap/buy-me-a-car/crawlers"
	"github.com/rattrap/buy-me-a-car/model"
)

func main() {
	urls, err := getUrls()
	if err != nil {
		log.Fatalf("Failed to parse pages: %s", err)
	}

	// tick := time.Tick(2 * time.Second)
	// for range tick {
	crawlUrls(urls)
	// }
}

func getUrls() ([]string, error) {
	file, err := os.Open("urls.ini")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	return urls, scanner.Err()
}

func crawlUrls(urls []string) {
	cars := []model.Car{}
	for _, u := range urls {
		u, err := url.Parse(u)
		if err != nil {
			log.Fatal(err)
		}

		if s.Contains(u.Host, "olx") {
			olx := crawlers.Crawl{crawlers.Olx{}}
			cars = append(olx.Crawl(u), cars...)
		} else if s.Contains(u.Host, "carzz") {
			carzz := crawlers.Crawl{crawlers.Carzz{}}
			cars = append(carzz.Crawl(u), cars...)
		} else if s.Contains(u.Host, "autovit") {
			autovit := crawlers.Crawl{crawlers.Autovit{}}
			cars = append(autovit.Crawl(u), cars...)
		} else {
			log.Fatalf("Crawler not found for: %s", u.Host)
		}
	}

	slcB, _ := json.Marshal(cars)
	log.Print(string(slcB))
}

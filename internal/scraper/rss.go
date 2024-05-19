package scraper

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"sync"
)

func FetchAllUrl(urls []string) error {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			_, err := FetchDataFromUrl(url)
			if err != nil {
				log.Println(err)
			}
		}(url)

	}
	wg.Wait()
	return nil
}

func FetchDataFromUrl(url string) (Rss, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return Rss{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return Rss{}, err
	}
	x := Rss{}
	err = xml.Unmarshal(body, &x)
	if err != nil {
		log.Println(err, url)
		return Rss{}, err
	}
	log.Println(x.Channel.Title)
	return x, nil
}

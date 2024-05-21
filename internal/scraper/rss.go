package scraper

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"sync"
)

func FetchAllUrl(urls []string) map[string]RssFeedXml {
	var wg sync.WaitGroup
	result := make(map[string]RssFeedXml, len(urls))
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			rss, err := FetchDataFromUrl(url)
			if err != nil {
				log.Println(err)
			}
			result[url] = rss
		}(url)
	}
	wg.Wait()
	return result
}

func FetchDataFromUrl(url string) (RssFeedXml, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return RssFeedXml{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return RssFeedXml{}, err
	}
	x := RssFeedXml{}
	err = xml.Unmarshal(body, &x)
	if err != nil {
		log.Println(err, url)
		return RssFeedXml{}, err
	}
	// TODO: make a table to store this
	// for _, item := range x.Channel.Items {
	// 	fmt.Println(item)
	// }
	return x, nil
}

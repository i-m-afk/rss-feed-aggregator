package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/i-m-afk/rss/internal/database"
	"github.com/i-m-afk/rss/internal/scraper"
)

// worker to fetch feeds continously
func worker(n int, conf *apiConfig, d time.Duration, done chan bool) {
	tick := 1
	ticker := time.NewTicker(d)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			fmt.Println(tick)
			tick++
			getFeedsToFetch(int32(n), conf)
		}
	}
}

func getFeedsToFetch(n int32, conf *apiConfig) {
	feeds, err := conf.DB.GetNextFeedsToFetch(context.Background(), n)
	if err != nil {
		log.Println(err)
		return
	}
	if len(feeds) == 0 {
		log.Println("no feeds to fetch")
		return
	}
	urls := make([]string, len(feeds))
	for i, feed := range feeds {
		urls[i] = feed.Url
	}
	if err := scraper.FetchAllUrl(urls); err == nil {
		setAllFeedsAsFetched(feeds, conf)
	} else {
		log.Println("unable to fetch all urls")
		return
	}
	fmt.Println("fetched these : ", urls)
}

func setAllFeedsAsFetched(feeds []database.Feed, conf *apiConfig) {
	for _, feed := range feeds {
		err := conf.DB.MarkFeedAsFetched(context.Background(), database.MarkFeedAsFetchedParams{
			LastFetchedAt: sql.NullTime{Time: time.Now()},
			ID:            feed.ID,
		})
		if err != nil {
			log.Println("Error making feed as fetched: ", err)
		}
	}
}

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/i-m-afk/rss/internal/database"
	"github.com/i-m-afk/rss/internal/scraper"
	"github.com/lib/pq"
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
	rssfeeds := scraper.FetchAllUrl(urls)
	conf.setAllFeedsAsFetched(feeds)
	conf.savePost(rssfeeds, feeds)
}

func (conf *apiConfig) setAllFeedsAsFetched(feeds []database.Feed) {
	for _, feed := range feeds {
		err := conf.DB.MarkFeedAsFetched(context.Background(), database.MarkFeedAsFetchedParams{
			LastFetchedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
			ID:            feed.ID,
		})
		if err != nil {
			log.Println("Error making feed as fetched: ", err)
		}
	}
}

func (conf *apiConfig) savePost(rss map[string]scraper.RssFeedXml, feeds []database.Feed) {
	for k, r := range rss {
		post, err := conf.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   sql.NullTime{Time: time.Now().UTC()},
			Title:       r.Channel.Title,
			Description: sql.NullString{String: r.Channel.Description, Valid: true},
			Url:         k,
			FeedID:      uuid.NullUUID{UUID: getFeedIdFromUrl(feeds, k), Valid: true},
			PublishedAt: sql.NullString{String: r.Channel.PubDate, Valid: true},
		})
		if err != nil {
			if perr, ok := err.(*pq.Error); ok && perr.Code.Name() == "unique_violation" {
				log.Println("Post already exists", r.Channel.Title)
				continue
			} else {
				log.Println(err)
			}
		} else {
			conf.saveRssItems(r.Channel.Items, post.ID)
		}
	}
}

func getFeedIdFromUrl(feeds []database.Feed, url string) uuid.UUID {
	for _, f := range feeds {
		if f.Url == url {
			return f.ID
		}
	}
	return uuid.Nil
}

func (conf *apiConfig) saveRssItems(items []*scraper.RssItem, postId uuid.UUID) {
	for _, item := range items {
		conf.saveRssItem(item, postId)
	}
}

func (conf *apiConfig) saveRssItem(item *scraper.RssItem, postId uuid.UUID) {
	_, err := conf.DB.CreateRssItem(context.Background(), database.CreateRssItemParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   sql.NullTime{Time: time.Now().UTC()},
		Title:       item.Title,
		Description: sql.NullString{String: item.Description, Valid: true},
		Url:         item.Link,
		Author:      sql.NullString{String: item.Author, Valid: true},
		PublishedAt: sql.NullString{String: item.PubDate, Valid: true},
		PostID:      uuid.NullUUID{UUID: postId, Valid: true},
	})
	if err != nil {
		if perr, ok := err.(*pq.Error); ok && perr.Code.Name() == "unique_violation" {
			log.Println("Rss Item already exists", item.Title)
			return
		}
		log.Println(err)
		return
	}
}

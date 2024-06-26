package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/i-m-afk/rss/internal/database"
)

type User struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
	Name      string    `json:"name"`
	ID        uuid.UUID `json:"id"`
}

type Feed struct {
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	LastFetchedAt time.Time `json:"last_fetched_at"`
	Url           string    `json:"url"`
	Name          string    `json:"name"`
	UserID        uuid.UUID `json:"user_id"`
	ID            uuid.UUID `json:"id"`
}

type FeedFollow struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        uuid.UUID `json:"id"`
	FeedId    uuid.UUID `json:"feed_id"`
	UserId    uuid.UUID `json:"user_id"`
}

type FeedAndFollows struct {
	Feed       Feed       `json:"feed"`
	FeedFollow FeedFollow `json:"feed_follow"`
}

type Post struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	PublishedAt string    `json:"published_at"`
	ID          uuid.UUID `json:"id"`
	FeedID      uuid.UUID `json:"feed_id"`
}

type RssItem struct {
	RssItemID          uuid.UUID `json:"id"`
	RssItemCreatedAt   time.Time `json:"created_at"`
	RssItemUpdatedAt   time.Time `json:"updated_at"`
	RssItemTitle       string    `json:"title"`
	RssItemUrl         string    `json:"url"`
	RssItemAuthor      string    `json:"author"`
	RssItemDescription string    `json:"description"`
	RssItemPublishedAt string    `json:"published_at"`
	RssItemPostID      uuid.UUID `json:"post_id"`
}

func (cfg *apiConfig) databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		ApiKey:    user.ApiKey,
	}
}

func (cfg *apiConfig) databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:            feed.ID,
		Name:          feed.Name,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.CreatedAt,
		Url:           feed.Url,
		UserID:        feed.UserID.UUID,
		LastFetchedAt: feed.LastFetchedAt.Time,
	}
}

func (cfg *apiConfig) databaseFeedsToFeeds(feeds []database.Feed) []Feed {
	feedList := make([]Feed, len(feeds))
	for i, feed := range feeds {
		feedList[i] = cfg.databaseFeedToFeed(feed)
	}
	return feedList
}

func (cfg *apiConfig) databaseFeedFollowToFeedFollow(feedfollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		UserId:    feedfollow.UserID.UUID,
		FeedId:    feedfollow.FeedID.UUID,
		ID:        feedfollow.ID,
		CreatedAt: feedfollow.CreatedAt,
		UpdatedAt: feedfollow.UpdatedAt,
	}
}

func (cfg *apiConfig) databaseFeedsFollowsToFeedsFollows(feedsfollows []database.FeedFollow) []FeedFollow {
	feedFollowList := make([]FeedFollow, len(feedsfollows))
	for i, feedfollow := range feedsfollows {
		feedFollowList[i] = cfg.databaseFeedFollowToFeedFollow(feedfollow)
	}
	return feedFollowList
}

func (cfg *apiConfig) databaseFeedNFollowsToFeedNFollows(feed database.Feed, feedfollow database.FeedFollow) FeedAndFollows {
	f := cfg.databaseFeedToFeed(feed)
	ff := cfg.databaseFeedFollowToFeedFollow(feedfollow)
	return FeedAndFollows{Feed: f, FeedFollow: ff}
}

func (cfg *apiConfig) databasePostToPost(post database.Post) Post {
	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt.Time,
		Title:       post.Title,
		Url:         post.Url,
		Description: post.Description.String,
		PublishedAt: post.PublishedAt.String,
		FeedID:      post.FeedID.UUID,
	}
}

func (cfg *apiConfig) databasePostsToPosts(posts []database.Post) []Post {
	postList := make([]Post, len(posts))
	for i, post := range posts {
		postList[i] = cfg.databasePostToPost(post)
	}
	return postList
}

func (cfg *apiConfig) databaseRssItemToRssItem(rssItem database.GetRssItemsForUserRow) RssItem {
	return RssItem{
		RssItemID:          rssItem.RssItemID,
		RssItemCreatedAt:   rssItem.RssItemCreatedAt,
		RssItemUpdatedAt:   rssItem.RssItemUpdatedAt.Time,
		RssItemTitle:       rssItem.RssItemTitle,
		RssItemUrl:         rssItem.RssItemUrl,
		RssItemAuthor:      rssItem.RssItemAuthor.String,
		RssItemDescription: rssItem.RssItemDescription.String,
		RssItemPublishedAt: rssItem.RssItemPublishedAt.String,
		RssItemPostID:      rssItem.RssItemPostID.UUID,
	}
}

func (cfg *apiConfig) databaseRssItemsToRssItems(rssItems []database.GetRssItemsForUserRow) []RssItem {
	rssItemList := make([]RssItem, len(rssItems))
	for i, rssItem := range rssItems {
		rssItemList[i] = cfg.databaseRssItemToRssItem(rssItem)
	}
	return rssItemList
}

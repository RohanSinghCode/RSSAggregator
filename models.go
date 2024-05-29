package main

import (
	"database/sql"
	"time"

	"github.com/RohanSinghCode/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Createdat: dbUser.Createdat,
		Updatedat: dbUser.Updatedat,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID            uuid.UUID
	Createdat     time.Time
	Updatedat     sql.NullTime
	Name          string
	Url           string
	UserID        uuid.UUID
	LastFetchedAt sql.NullTime
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:            dbFeed.ID,
		Createdat:     dbFeed.Createdat,
		Updatedat:     dbFeed.Updatedat,
		Name:          dbFeed.Name,
		Url:           dbFeed.Url,
		UserID:        dbFeed.UserID,
		LastFetchedAt: dbFeed.LastFetchedAt,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}

	return feeds
}

type FeedFollows struct {
	ID        uuid.UUID
	Createdat time.Time
	Updatedat sql.NullTime
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

func databaseFeedFollowsToFeedFollows(dbFeed database.FeedFollow) FeedFollows {
	return FeedFollows{
		ID:        dbFeed.ID,
		Createdat: dbFeed.Createdat,
		Updatedat: dbFeed.Updatedat,
		UserID:    dbFeed.UserID,
		FeedID:    dbFeed.FeedID,
	}
}

func databaseFeedFollowersToFeedFollowers(dbFeedFollows []database.FeedFollow) []FeedFollows {
	feeds := []FeedFollows{}
	for _, dbFeedFollow := range dbFeedFollows {
		feeds = append(feeds, databaseFeedFollowsToFeedFollows(dbFeedFollow))
	}

	return feeds
}

type Post struct {
	ID          uuid.UUID
	Createdat   time.Time
	Updatedat   time.Time
	Title       string
	Description *string
	PublishedAt time.Time
	URL         string
	FeedID      uuid.UUID
}

func databasePostToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}
	return Post{
		ID:          dbPost.ID,
		Createdat:   dbPost.CreatedAt,
		Updatedat:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		URL:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, dbPost := range dbPosts {
		posts = append(posts, databasePostToPost(dbPost))
	}

	return posts
}

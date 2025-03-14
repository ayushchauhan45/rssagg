package main

import (
	"time"
	"github.com/ayushchauhan_45/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Name      string     `json:"name"`
	APIKey    string      `json:"api_key"`
}

func  databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID   `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Name      string      `json:"name"`
	Url       string      `json:"url"`
	UserID    uuid.UUID   `json:"user_id"`
}

func databaseFeedtoFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedstoFeeds(dbFeed []database.Feed) []Feed {
	feeds := []Feed{}
	for _,dbFeeds := range dbFeed {
		feeds = append(feeds, databaseFeedtoFeed(dbFeeds))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	UserID    uuid.UUID  `json:"user_id"`
	FeedID    uuid.UUID   `json:"feed_id"`  
}


func databaseFeedFollowtoFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}
}

func databaseFeedFollowstoFeedFollows(dbFeedFollow []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _,dbFeedFollows := range dbFeedFollow {
		feedFollows = append(feedFollows, databaseFeedFollowtoFeedFollow(dbFeedFollows))
	}
	return feedFollows
}

type Post struct{
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	Descreption *string `json:"description"`
	PublishedAt time.Time    `json:"published_at"`
	Url         string      `json:"url"`
	FeedID      uuid.UUID    `json:"feed_id"`
}

func databasePosttoPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Descreption.Valid{
		description = &dbPost.Descreption.String
	}
	return Post{
		ID:		  dbPost.ID,
		CreatedAt: dbPost.CreatedAt,
		UpdatedAt: dbPost.UpdatedAt,
		Title:     dbPost.Title,
		Descreption: description,
		PublishedAt: dbPost.PublishedAt,
		Url:       dbPost.Url,
		FeedID:    dbPost.FeedID,
	}
}

func databasePoststoPosts(dbPost []database.Post) []Post {
	posts := []Post{}
	for _,dbPosts := range dbPost {
		posts = append(posts, databasePosttoPost(dbPosts))
	}
	return posts
}
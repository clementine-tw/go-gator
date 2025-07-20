package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/clementine-tw/go-gator/internal/database"
)

func handlerAggregate(s *state, c command) error {

	if len(c.Args) == 0 {
		return fmt.Errorf("usage: %s <fetch_interval>", c.Name)
	}

	fetch_interval, err := time.ParseDuration(c.Args[0])
	if err != nil {
		return fmt.Errorf("parse time interval string error: %w", err)
	}

	log.Printf("collect feeds every %s...", fetch_interval)

	ticker := time.NewTicker(fetch_interval)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("couldn't get next feed to fetch: %v", err)
		return
	}

	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(
		context.Background(),
		feed.ID,
	)
	if err != nil {
		log.Printf("couldn't mark feed %s as fetched: %v", feed.Name, err)
		return
	}

	content, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("couldn't fetch feed %s: %v", feed.Name, err)
		return
	}

	for _, item := range content.Channel.Item {
		log.Printf("Found post: %s\n", item.Title)
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(content.Channel.Item))
}

func handlerAddFeed(s *state, c command, user database.User) error {

	if len(c.Args) < 2 {
		return fmt.Errorf("usage: %s <feed_name> <feed_url>", c.Name)
	}

	feedName := c.Args[0]
	feedURL := c.Args[1]

	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			Name:   feedName,
			Url:    feedURL,
			UserID: user.ID,
		})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			UserID: user.ID,
			FeedID: feed.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	fmt.Println("feed created successfully:")
	fmt.Printf("%+v\n", feed)

	return nil
}

func handlerFeeds(s *state, _ command) error {

	feeds, err := s.db.GetFeedsWithUserName(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	fmt.Println("Feeds:\n-")
	for _, feed := range feeds {
		fmt.Printf("name: %s\nurl: %s\nuser: %s\n-\n", feed.Name, feed.Url, feed.UserName)
	}
	return nil
}

package main

import (
	"context"
	"fmt"

	"github.com/clementine-tw/go-gator/internal/database"
)

func handlerAggregate(_ *state, _ command) error {

	url := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't fetch RSS feed: %w", err)
	}

	fmt.Printf("%v\n", feed)
	return nil
}

func handlerAddFeed(s *state, c command) error {

	if len(c.Args) < 2 {
		return fmt.Errorf("usage: %s <feed_name> <feed_url>", c.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
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

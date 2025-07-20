package main

import (
	"context"
	"fmt"

	"github.com/clementine-tw/go-gator/internal/database"
)

func handlerFollowFeed(s *state, c command) error {

	if len(c.Args) == 0 {
		return fmt.Errorf("usage: %s <feed_url>", c.Name)
	}
	feedUrl := c.Args[0]

	feedID, err := s.db.GetFeedIDByUrl(
		context.Background(),
		feedUrl,
	)
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	user, err := s.db.GetUser(
		context.Background(),
		s.cfg.CurrentUserName,
	)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	record, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			FeedID: feedID,
			UserID: user.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	fmt.Println("Follow feed successfully:")
	fmt.Printf(" * Feed Name: %s\n", record.FeedName)

	return nil
}

func handlerListFeedFollows(s *state, _ command) error {

	user, err := s.db.GetUser(
		context.Background(),
		s.cfg.CurrentUserName,
	)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	feeds, err := s.db.GetFollowingFeedsByUserID(
		context.Background(),
		user.ID,
	)

	if len(feeds) == 0 {
		fmt.Println("No following feed")
		return nil
	}

	fmt.Println("Following feeds:")
	for _, feed := range feeds {
		fmt.Printf(" * %s\n", feed.FeedName)
	}

	return nil
}

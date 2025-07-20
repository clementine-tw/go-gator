package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/clementine-tw/go-gator/internal/database"
)

func handlerBrowse(s *state, c command, user database.User) error {
	limit := 2
	if len(c.Args) > 0 {
		i, err := strconv.Atoi(c.Args[0])
		if err != nil {
			return fmt.Errorf("usage: %s [limit_number]", c.Name)
		}
		limit = i
	}

	posts, err := s.db.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  int32(limit),
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	fmt.Printf("Found %v posts\n\n", len(posts))
	for _, post := range posts {
		fmt.Printf("%v from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf(" === %s === \n", post.Title)
		fmt.Println(post.Description.String)
		fmt.Printf("Link: %s\n\n", post.Url)
	}

	return nil
}

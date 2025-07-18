package main

import (
	"context"
	"fmt"
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

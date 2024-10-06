package main

import (
	"context"
	"fmt"
	"gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	ctx := context.Background()

	feed, err := rss.FetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}

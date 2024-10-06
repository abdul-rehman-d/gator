package main

import (
	"context"
	"fmt"
	"gator/internal/utils"
)

func handlerGetAllFeeds(s *state, cmd command) error {
	ctx := context.Background()
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("no users exist.")
		}
		return err
	}
	for _, feed := range feeds {
		fmt.Printf("Name:	%s\n", feed.Feed.Name)
		fmt.Printf("URL:	%s\n", feed.Feed.Url)
		fmt.Printf("User:	%s\n", feed.User.Name)
	}
	return nil
}

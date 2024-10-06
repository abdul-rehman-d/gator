package main

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/database"
	"gator/internal/utils"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("Usage: gator addFeed <title> <url>")
	}
	ctx := context.Background()

	user, err := s.db.GetUser(ctx, s.config.CurrentUsername)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("%s does not exist. Try registering first", s.config.CurrentUsername)
		}
		return err
	}
	currentTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}

	insertedFeed, err := s.db.CreateFeed(ctx, feed)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", insertedFeed)
	return nil
}

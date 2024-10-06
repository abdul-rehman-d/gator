package main

import (
	"context"
	"fmt"
	"gator/internal/utils"
)

func handlerGetUsers(s *state, cmd command) error {
	ctx := context.Background()
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("no users exist.")
		}
		return err
	}
	for _, user := range users {
		if user.Name == s.config.CurrentUsername {
			fmt.Printf("%s (current)\n", user.Name)
		} else {
			fmt.Printf("%s\n", user.Name)
		}
	}
	return nil
}

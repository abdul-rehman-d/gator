package main

import (
	"context"
	"fmt"
	"gator/internal/utils"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Usage: gator login <username>")
	}
	username := cmd.args[0]

	ctx := context.Background()
	_, err := s.db.GetUser(ctx, username)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("%s does not exist. Try registering first", username)
		}
		return err
	}

	err = s.config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Println("User has been set.")
	return nil
}

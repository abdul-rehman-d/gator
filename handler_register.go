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

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Usage: gator login <username>")
	}
	username := cmd.args[0]
	ctx := context.Background()

	currentTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	user := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      username,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	insertedUser, err := s.db.CreateUser(ctx, user)
	if err != nil {
		if utils.IsDuplicateError(err) {
			return fmt.Errorf("%s already exists. Try a different username.", username)
		}
		return err
	}
	fmt.Printf("%v\n", insertedUser)
	err = s.config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Println("User has been set.")
	return nil
}

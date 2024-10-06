package main

import "context"

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	return s.db.ResetAllUsers(ctx)
}

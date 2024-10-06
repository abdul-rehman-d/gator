package main

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"gator/internal/utils"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if f := c.commands[cmd.name]; f != nil {
		return f(s, cmd)
	}
	return fmt.Errorf(`error: %v is not a command.`, cmd.name)
}

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

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	return s.db.ResetAllUsers(ctx)
}

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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no command given.")
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)

	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	s := &state{
		db:     dbQueries,
		config: &cfg,
	}

	cmds := &commands{
		commands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)

	if err := cmds.run(s, cmd); err != nil {
		log.Fatal(err)
	}
}

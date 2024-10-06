package main

import (
	"fmt"
	"gator/internal/config"
	"log"
	"os"
)

type state struct {
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
	err := s.config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Println("User has been set.")
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

	s := &state{
		config: &cfg,
	}

	cmds := &commands{
		commands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if err := cmds.run(s, cmd); err != nil {
		log.Fatal(err)
	}
}

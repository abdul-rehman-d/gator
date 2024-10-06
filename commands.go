package main

import "fmt"

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

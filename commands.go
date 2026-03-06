package main

import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	hander, ok := c.handlers[cmd.name]
	if !ok {
		return errors.New("command doesn't exist")
	}
	err := hander(s, cmd)
	return err

}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

package commands

import (
	"errors"

	"github.com/brettsvoid/gator/internal/config"
	"github.com/brettsvoid/gator/internal/database"
)

type State struct {
	DB     *database.Queries
	Config *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Handlers map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, ok := c.Handlers[cmd.Name]
	if !ok {
		return errors.New("handler for command does not exist")
	}

	err := handler(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Handlers[name] = f
}

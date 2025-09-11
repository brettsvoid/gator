package middleware

import (
	"context"
	"fmt"

	"github.com/brettsvoid/gator/internal/commands"
	"github.com/brettsvoid/gator/internal/database"
)

func LoggedIn(
	handler func(s *commands.State, cmd commands.Command, user database.User) error,
) func(*commands.State, commands.Command) error {
	return func(s *commands.State, cmd commands.Command) error {
		user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("couldn't find user: %w", err)
		}

		return handler(s, cmd, user)
	}
}

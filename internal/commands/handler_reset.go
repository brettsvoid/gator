package commands

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {
	err := s.DB.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}

	fmt.Println("Database has been reset successfully!")
	return nil
}

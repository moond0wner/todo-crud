package users_service

import (
	"context"
	"fmt"

	core_errors "github.com/moond0wner/todo-nilchan/internal/core/errors"
)

func (s *UsersService) DeleteUser(
	ctx context.Context,
	userID int,
) error {
	if userID <= 0 {
		return fmt.Errorf(
			"userID must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	err := s.usersRepository.DeleteUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("delete user from repository: %w", err)
	}
	return nil
}

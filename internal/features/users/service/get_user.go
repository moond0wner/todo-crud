package users_service

import (
	"context"
	"fmt"

	"github.com/moond0wner/todo-nilchan/internal/core/domain"
	core_errors "github.com/moond0wner/todo-nilchan/internal/core/errors"
)

func (s *UsersService) GetUser(
	ctx context.Context,
	userID int,
) (domain.User, error) {
	if userID <= 0 {
		return domain.User{}, fmt.Errorf(
			"userID must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	user, err := s.usersRepository.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	return user, nil
}

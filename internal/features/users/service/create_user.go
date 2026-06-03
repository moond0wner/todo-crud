package users_service

import (
	"context"
	"fmt"

	"github.com/moond0wner/todo-nilchan/internal/core/domain"
)

func NewUsersService(us UsersRepository) UsersService {
	return UsersService{
		usersRepository: us,
	}
}

func (s *UsersService) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domain: %w", err)
	}

	user, err := s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}

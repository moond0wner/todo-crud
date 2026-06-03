package users_service

import (
	"context"
	"fmt"

	"github.com/moond0wner/todo-nilchan/internal/core/domain"
	core_errors "github.com/moond0wner/todo-nilchan/internal/core/errors"
)

func (s *UsersService) PatchUser(
	ctx context.Context,
	id int,
	patch domain.UserPatch,
) (domain.User, error) {
	if id <= 0 {
		return domain.User{}, fmt.Errorf(
			"id must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	user, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("apply user patch: %w", err)
	}

	patchedUser, err := s.usersRepository.PatchUser(
		ctx,
		id,
		user,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("patch user in repository: %w", err)
	}

	return patchedUser, nil
}

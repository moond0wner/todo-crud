package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/moond0wner/todo-nilchan/internal/core/domain"
	core_errors "github.com/moond0wner/todo-nilchan/internal/core/errors"
	core_postgres_pool "github.com/moond0wner/todo-nilchan/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) GetUser(
	ctx context.Context,
	userID int,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	sqlQuery := `
	SELECT id, version, full_name, phone_number 
	FROM todoapp.users
	WHERE id=$1;
	`

	row := r.pool.QueryRow(
		ctx,
		sqlQuery,
		userID,
	)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id='%d': %w",
				userID,
				core_errors.ErrNotFound,
			)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := userDomainFromModel(userModel)

	return userDomain, nil

}

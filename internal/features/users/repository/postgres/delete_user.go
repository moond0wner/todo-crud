package users_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/moond0wner/todo-nilchan/internal/core/errors"
)

func (r *UsersRepository) DeleteUser(
	ctx context.Context,
	userID int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	sqlQuery := `
	DELETE 
	FROM todoapp.users
	WHERE id=$1;
	`
	cmdTag, err := r.pool.Exec(
		ctx,
		sqlQuery,
		userID,
	)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user with id='%d': %w", userID, core_errors.ErrNotFound)
	}

	return nil
}

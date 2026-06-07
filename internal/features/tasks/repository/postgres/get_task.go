package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/moond0wner/todo-nilchan/internal/core/domain"
	core_errors "github.com/moond0wner/todo-nilchan/internal/core/errors"
	core_postgres_pool "github.com/moond0wner/todo-nilchan/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) GetTask(
	ctx context.Context,
	taskID int,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	sqlQuery := `
	SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
	FROM todoapp.tasks
	WHERE id=$1;
	`

	row := r.pool.QueryRow(
		ctx,
		sqlQuery,
		taskID,
	)

	var taskModel TaskModel
	if err := row.Scan(
		&taskModel.ID,
		&taskModel.Version,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CreatedAt,
		&taskModel.CompletedAt,
		&taskModel.AuthorUserID,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"task with id='%d': %w",
				taskID,
				core_errors.ErrNotFound,
			)
		}
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	taskDomain := taskDomainFromModel(taskModel)
	return taskDomain, nil
}

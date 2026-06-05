package tasks_postgres_repository

import (
	"context"
	"fmt"

	"github.com/moond0wner/todo-nilchan/internal/core/domain"
)

func (r *TasksRepository) GetTasks(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	sqlQuery := `
	SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
	FROM todoapp.tasks
	%s
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`
	args := []any{limit, offset}

	if userID != nil {
		sqlQuery = fmt.Sprintf(sqlQuery, "WHERE author_user_id=$3")
		args = append(args, userID)
	} else {
		sqlQuery = fmt.Sprintf(sqlQuery, "")
	}

	rows, err := r.pool.Query(
		ctx,
		sqlQuery,
		args...,
	)
	if err != nil {
		return nil, fmt.Errorf("select users: %w", err)
	}

	defer rows.Close()

	var tasks []TaskModel
	for rows.Next() {
		var taskModel TaskModel
		if err := rows.Scan(
			&taskModel.ID,
			&taskModel.Version,
			&taskModel.Title,
			&taskModel.Description,
			&taskModel.Completed,
			&taskModel.CreatedAt,
			&taskModel.CompletedAt,
			&taskModel.AuthorUserID,
		); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		tasks = append(tasks, taskModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	tasksDomain := tasksDomainFromModels(tasks)
	return tasksDomain, nil
}

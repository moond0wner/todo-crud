package tasks_transport_http

import (
	"time"

	"github.com/moond0wner/todo-nilchan/internal/core/domain"
)

type TaskDTOResponse struct {
	ID           int        `json:"id" example:"15"`
	Version      int        `json:"version" example:"1"`
	Title        string     `json:"title" example:"Домашнее задание"`
	Description  *string    `json:"description" example:"Выполнить задание до четверга"`
	Completed    bool       `json:"completed" example:"false"`
	CreatedAt    time.Time  `json:"created_at" example:"2026-02-26T10:30:00Z"`
	CompletedAt  *time.Time `json:"completed_at" example:"null"`
	AuthorUserID int        `json:"author_user_id" example:"4"`
}

func taskDTOFromDomain(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}

func tasksDTOFromDomains(tasks []domain.Task) []TaskDTOResponse {
	tasksDTO := make([]TaskDTOResponse, len(tasks))
	for i, task := range tasks {
		tasksDTO[i] = taskDTOFromDomain(task)
	}
	return tasksDTO
}

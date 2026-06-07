package statistics_service

import (
	"context"
	"time"

	"github.com/moond0wner/todo-nilchan/internal/core/domain"
)

type StatisticsService struct {
	statisticsRepository StatisticsRepository
}

type StatisticsRepository interface {
	GetTasks(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) ([]domain.Task, error)
}

func NewStatisticsService(sr StatisticsRepository) *StatisticsService {
	return &StatisticsService{
		statisticsRepository: sr,
	}
}

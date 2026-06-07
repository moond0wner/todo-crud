package statistics_transport_http

import (
	"context"
	"time"

	"github.com/moond0wner/todo-nilchan/internal/core/domain"
	core_http_server "github.com/moond0wner/todo-nilchan/internal/core/transport/http/server"
)

type StatiscticsHTTPHandler struct {
	statisticsService StatisticsService
}

type StatisticsService interface {
	GetStatistics(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(ss StatisticsService) *StatiscticsHTTPHandler {
	return &StatiscticsHTTPHandler{
		statisticsService: ss,
	}
}

func (h *StatiscticsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  "GET",
			Path:    "/statistics",
			Handler: h.GetStatistics,
		},
	}
}

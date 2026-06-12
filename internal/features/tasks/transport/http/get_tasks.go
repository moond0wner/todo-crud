package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/moond0wner/todo-nilchan/internal/core/logger"
	core_http_request "github.com/moond0wner/todo-nilchan/internal/core/transport/http/request"
	core_http_response "github.com/moond0wner/todo-nilchan/internal/core/transport/http/response"
)

type GetTasksResponse []TaskDTOResponse

// GetTasks 	godoc
// @Summary 	Список задач
// @Description Просмотр всех задач с опциональной пагинацией и/или фильтрацией по ID автора задач
// @Tags 		tasks
// @Produce		json
// @Param		user_id query	int  		false			  "Фильтрая задач по ID пользователя"
// @Param 		limit  	query 	int  		false 			  "Размер страницы с задачами"
// @Param 		offset 	query 	int  		false 			  "Смещение страницы с задачами"
// @Success 	200 {object} GetTasksResponse 				  "Успешное получение списка задач"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal servver error"
// @Router		/tasks [get]
func (h *TasksHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	queryParams, err := getUserIDLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'user_id'/'limit'/'offset' query param",
		)
		return
	}

	tasks, err := h.tasksService.GetTasks(
		ctx,
		queryParams.userID,
		queryParams.limit,
		queryParams.offset,
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get /tasks",
		)
		return
	}

	response := GetTasksResponse(tasksDTOFromDomains(tasks))
	responseHandler.JSONResponse(
		response,
		http.StatusOK,
	)
}

type queryParams struct {
	userID *int
	limit  *int
	offset *int
}

func getUserIDLimitOffsetQueryParams(r *http.Request) (queryParams, error) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
		UserIDKey           = "user_id"
	)

	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return queryParams{}, fmt.Errorf("get 'limit' query param: %w:", err)
	}
	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return queryParams{}, fmt.Errorf("get 'offset' query param: %w", err)
	}
	userID, err := core_http_request.GetIntQueryParam(r, UserIDKey)
	if err != nil {
		return queryParams{}, fmt.Errorf("get `user_id` query param: %w", err)
	}

	return queryParams{
		userID: userID,
		limit:  limit,
		offset: offset,
	}, nil
}

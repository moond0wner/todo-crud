package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/moond0wner/todo-nilchan/internal/core/logger"
	core_http_request "github.com/moond0wner/todo-nilchan/internal/core/transport/http/request"
	core_http_response "github.com/moond0wner/todo-nilchan/internal/core/transport/http/response"
)

type GetTaskResponse TaskDTOResponse

// GetTask 		godoc
// @Summary 	Получение задачи
// @Description Получение конкретной задачи по ID
// @Tags 		tasks
// @Produce 	json
// @Param 		id 	path 	 int 			true 			  "ID получаемой задачи"
// @Success 	200 {object} GetTaskResponse 				  "Задача успешно найдена"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	404 {object} core_http_response.ErrorResponse "Task not found"
// @Failure 	400 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/tasks/{id} [get]
func (h *TasksHTTPHandler) GetTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get path param 'id'",
		)
		return
	}

	taskDomain, err := h.tasksService.GetTask(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get /tasks/{id}",
		)
		return
	}

	response := GetTaskResponse(taskDTOFromDomain(taskDomain))
	responseHandler.JSONResponse(
		response,
		http.StatusOK,
	)

}

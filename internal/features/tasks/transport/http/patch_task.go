package tasks_transport_http

import (
	"fmt"
	"net/http"

	"github.com/moond0wner/todo-nilchan/internal/core/domain"
	core_logger "github.com/moond0wner/todo-nilchan/internal/core/logger"
	core_http_request "github.com/moond0wner/todo-nilchan/internal/core/transport/http/request"
	core_http_response "github.com/moond0wner/todo-nilchan/internal/core/transport/http/response"
	core_http_types "github.com/moond0wner/todo-nilchan/internal/core/transport/http/types"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title" swaggertype:"string" example:"Встать на учебу"`
	Description core_http_types.Nullable[string] `json:"description" swaggertype:"string" example:"Проснуться утром в 6"`
	Completed   core_http_types.Nullable[bool]   `json:"completed" swaggertype:"boolean"`
}

type PatchTaskResponse TaskDTOResponse

// PatchUser 	godoc
// @Summary 	Обновить задачу
// @Description Обновление данных об уже существующей в системе задачи
// @Description ### Логика обновления полей (Three-state logic):
// @Description 1. **Поле не передано**: `description` игнорируется, значение в БД не меняется
// @Description 2. **Явно передано значение** : `"description": "Утром в 6:00 встать на учебу"` - устанавливает новое описание в БД
// @Description 3. **Передан null**: `"description": null` - очищает поле в БД (set to NULL)
// @Description Ограничения: `title` и `completed` не может быть выставлены как null
// @Tags 		tasks
// @Accept 		json
// @Produce 	json
// @Param 		id 		path int 				  true 		  "ID изменяемой задачи"
// @Param 		request body PatchTaskRequest     true 		  "PatchTask тело запроса"
// @Success 	200 {object} PatchTaskResponse 				  "Успешно измененная задача"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	404 {object} core_http_response.ErrorResponse "Task not found"
// @Failure 	409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/tasks/{id} [patch]
func (r *PatchTaskRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("`Title` can`t be NULL")
		}
		titlelength := len([]rune(*r.Title.Value))
		if titlelength < 1 || titlelength > 100 {
			return fmt.Errorf("`Title` must be between 1 and 100")
		}
	}

	if r.Description.Set && r.Description.Value != nil {
		descriptionLength := len([]rune(*r.Description.Value))
		if descriptionLength < 1 || descriptionLength > 1000 {
			return fmt.Errorf("`Description` must be between 1 and 1000")
		}
	}

	if r.Completed.Set {
		if r.Completed.Value == nil {
			return fmt.Errorf("`Completed` can`t be NULL")
		}
	}

	return nil
}

func (h *TasksHTTPHandler) PatchTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get path param 'id'",
		)
		return
	}

	var request PatchTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}

	taskPatch := taskPatchFromRequest(request)

	taskDomain, err := h.tasksService.PatchTask(
		ctx,
		taskID,
		taskPatch,
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch task",
		)
		return
	}

	response := PatchTaskResponse(taskDTOFromDomain(taskDomain))
	responseHandler.JSONResponse(
		response,
		http.StatusOK,
	)
}

func taskPatchFromRequest(request PatchTaskRequest) domain.TaskPatch {
	return domain.NewTaskPatch(
		request.Title.ToDomain(),
		request.Description.ToDomain(),
		request.Completed.ToDomain(),
	)
}

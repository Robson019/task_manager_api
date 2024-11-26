package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"task_manager/src/api/handlers/dto/request"
	"task_manager/src/api/handlers/dto/response"
	"task_manager/src/api/handlers/utils/converters"
	"task_manager/src/api/handlers/utils/params"
	"task_manager/src/core/domain/task"
	"task_manager/src/core/interfaces/primary"
	"task_manager/src/core/messages"
)

type TaskHandlers struct {
	service primary.TaskManager
}

// CreateTask
// @ID CreateTask
// @Summary Criar tarefa
// @Tags Tarefas
// @Description Rota que permite a criação de uma tarefa pelo usuário.
// @Security bearerAuth
// @Param json body request.TaskDTO true "JSON com todos os dados necessários para criar uma tarefa."
// @Accept json
// @Produce json
// @Success 201 {object} response.DefaultPostResponse "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /tasks [post]
func (instance TaskHandlers) CreateTask(ctx echo.Context) error {
	var taskDTO request.TaskDTO
	bindError := ctx.Bind(&taskDTO)
	if bindError != nil {
		return getDefaultBadRequestResponse(ctx)
	}

	_task, err := task.NewBuilder().WithTitle(taskDTO.Title).
		WithStatus(taskDTO.Status).WithDescription(taskDTO.Description).Build()
	if err != nil {
		return getHttpHandledErrorResponse(ctx, err)
	}

	id, err := instance.service.CreateTask(*_task)
	if err != nil {
		return getHttpHandledErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, response.DefaultPostResponse{ID: *id})
}

// FindTasks
// @ID FindTasks
// @Summary Listar tarefas
// @Tags Tarefas
// @Description Rota que permite a listagem das tarefas.
// @Security bearerAuth
// @Accept json
// @Produce json
// @Success 200 {array} response.Task "Requisição realizada com sucesso."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /tasks [get]
func (instance TaskHandlers) FindTasks(ctx echo.Context) error {
	taskRows, err := instance.service.FindTasks()
	if err != nil {
		return getHttpHandledErrorResponse(ctx, err)
	}

	tasks := make([]response.Task, 0)
	for _, row := range taskRows {
		_task := response.NewTask(row)
		tasks = append(tasks, *_task)
	}

	return ctx.JSON(http.StatusOK, tasks)
}

// FindTaskByID
// @ID FindTaskByID
// @Summary Buscar tarefa
// @Tags Tarefas
// @Description Rota que permite a busca de uma tarefa pelo id.
// @Security bearerAuth
// @Param id path string true "ID da tarefa." default(03b3aecd-1b52-4357-875c-298a4bc60132)
// @Accept json
// @Produce json
// @Success 200 {object} response.Task "Requisição realizada com sucesso."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /tasks/{id} [get]
func (instance TaskHandlers) FindTaskByID(ctx echo.Context) error {
	id, err := converters.ConvertFromStringToUUID(ctx.Param(params.ID),
		messages.TaskID, messages.TaskIDErrorMessage, messages.ConversionErrorMessage)
	if err != nil {
		return getHttpHandledErrorResponse(ctx, err)
	}

	_task, err := instance.service.FindTaskByID(*id)
	if err != nil {
		return getHttpHandledErrorResponse(ctx, err)
	}
	taskResponse := response.NewTask(*_task)

	return ctx.JSON(http.StatusOK, taskResponse)
}

// UpdateTask
// @ID UpdateTask
// @Summary Atualizar tarefa
// @Tags Tarefas
// @Description Rota que permite a atualização de uma tarefa pelo usuário.
// @Security bearerAuth
// @Accept json
// @Param id path string true "ID da tarefa." default(03b3aecd-1b52-4357-875c-298a4bc60132)
// @Param json body request.TaskDTO true "JSON com todos os dados necessários para criar uma tarefa."
// @Produce json
// @Success 204 "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /tasks/{id} [put]
func (instance TaskHandlers) UpdateTask(ctx echo.Context) error {
	id, err := converters.ConvertFromStringToUUID(ctx.Param(params.ID),
		messages.TaskID, messages.TaskIDErrorMessage, messages.ConversionErrorMessage)
	if err != nil {
		return getHttpHandledErrorResponse(ctx, err)
	}

	var taskDTO request.TaskDTO
	bindError := ctx.Bind(&taskDTO)
	if bindError != nil {
		return getDefaultBadRequestResponse(ctx)
	}

	_task, err := task.NewBuilder().WithID(*id).WithStatus(taskDTO.Status).
		WithTitle(taskDTO.Title).WithDescription(taskDTO.Description).Build()
	if err != nil {
		return getHttpHandledErrorResponse(ctx, err)
	}

	err = instance.service.UpdateTask(*_task)
	if err != nil {
		return getHttpHandledErrorResponse(ctx, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// DeleteTask
// @ID DeleteTask
// @Summary Deletar tarefa
// @Tags Tarefas
// @Description Rota que permite apagar uma tarefa.
// @Security bearerAuth
// @Accept json
// @Param id path string true "ID da tarefa." default(03b3aecd-1b52-4357-875c-298a4bc60132)
// @Produce json
// @Success 204 "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /tasks/{id} [delete]
func (instance TaskHandlers) DeleteTask(ctx echo.Context) error {
	id, err := converters.ConvertFromStringToUUID(ctx.Param(params.ID),
		messages.TaskID, messages.TaskIDErrorMessage, messages.ConversionErrorMessage)
	if err != nil {
		return getHttpHandledErrorResponse(ctx, err)
	}

	err = instance.service.DeleteTask(*id)
	if err != nil {
		return getHttpHandledErrorResponse(ctx, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func NewTaskHandlers(service primary.TaskManager) *TaskHandlers {
	return &TaskHandlers{service}
}

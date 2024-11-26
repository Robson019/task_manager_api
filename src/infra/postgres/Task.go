package postgres

import (
	"context"
	"github.com/google/uuid"
	"task_manager/src/core/domain/task"
	"task_manager/src/core/errors"
	"task_manager/src/core/interfaces/repository"
	"task_manager/src/core/messages"
	"task_manager/src/infra/postgres/bridge"
)

var _ repository.TaskLoader = &TaskPostgresRepository{}

type TaskPostgresRepository struct {
	connectorManager
}

func (instance TaskPostgresRepository) CreateTask(_task task.Task) (*uuid.UUID, errors.Error) {
	conn, err := instance.getConnection()
	if err != nil {
		return nil, errors.NewUnavailableServiceError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer instance.closeConnection(conn)

	query := bridge.New(conn)
	params := bridge.InsertIntoTaskParams{
		Title:       _task.Title(),
		Description: stringToNullString(_task.Description()),
		Status:      _task.Status(),
	}
	id, queryError := query.InsertIntoTask(context.Background(), params)
	if queryError != nil {
		return nil, errors.NewUnexpectedError(messages.InsertingDataErrorMessage, queryError)
	}

	return &id, nil
}

func (instance TaskPostgresRepository) FindTasks() ([]task.Task, errors.Error) {
	conn, err := instance.getConnection()
	if err != nil {
		return nil, errors.NewUnavailableServiceError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer instance.closeConnection(conn)

	query := bridge.New(conn)
	tasksRow, err := query.SelectTasks(context.Background())
	if err != nil {
		return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, err)
	}

	tasks := make([]task.Task, 0)
	for _, t := range tasksRow {
		taskBuilder := task.NewBuilder()
		taskBuilder.WithID(t.TaskID).WithTitle(t.TaskTitle).WithDescription(t.TaskDescription.String).
			WithStatus(t.TaskStatus).WithCreatedAt(t.TaskCreatedAt.Time).WithUpdatedAt(t.TaskUpdatedAt.Time)

		_task, validationError := taskBuilder.Build()
		if validationError != nil {
			return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, validationError)
		}

		tasks = append(tasks, *_task)
	}

	return tasks, nil
}

func (instance TaskPostgresRepository) FindTaskByID(taskID uuid.UUID) (*task.Task, errors.Error) {
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return nil, errors.NewUnavailableServiceError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer instance.closeConnection(conn)

	query := bridge.New(conn)
	row, bedError := query.SelectTaskByID(context.Background(), taskID)
	if bedError != nil {
		return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, bedError)
	}

	taskBuilder := task.NewBuilder()
	taskBuilder.WithID(row.TaskID).WithTitle(row.TaskTitle).WithDescription(row.TaskDescription.String).
		WithStatus(row.TaskStatus).WithCreatedAt(row.TaskCreatedAt.Time).WithUpdatedAt(row.TaskUpdatedAt.Time)

	_task, validationError := taskBuilder.Build()

	if validationError != nil {
		return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, validationError)
	}

	return _task, nil
}

func (instance TaskPostgresRepository) UpdateTask(_task task.Task) errors.Error {
	conn, err := instance.getConnection()
	if err != nil {
		return errors.NewUnavailableServiceError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer instance.closeConnection(conn)

	query := bridge.New(conn)
	params := bridge.UpdateTaskParams{
		TaskTitle:       _task.Title(),
		TaskDescription: stringToNullString(_task.Description()),
		TaskStatus:      _task.Status(),
		TaskID:          _task.ID(),
	}
	err = query.UpdateTask(context.Background(), params)
	if err != nil {
		return errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
	}

	return nil
}

func (instance TaskPostgresRepository) DeleteTask(taskID uuid.UUID) errors.Error {
	conn, err := instance.getConnection()
	if err != nil {
		return errors.NewUnavailableServiceError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer instance.closeConnection(conn)

	query := bridge.New(conn)
	err = query.DeleteTask(context.Background(), taskID)
	if err != nil {
		return errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
	}

	return nil
}

func NewTaskPostgresRepository(manager connectorManager) *TaskPostgresRepository {
	return &TaskPostgresRepository{manager}
}

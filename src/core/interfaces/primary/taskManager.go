package primary

import (
	"github.com/google/uuid"
	"task_manager/src/core/domain/task"
	"task_manager/src/core/errors"
)

type TaskManager interface {
	CreateTask(_task task.Task) (*uuid.UUID, errors.Error)
	FindTasks() ([]task.Task, errors.Error)
	FindTaskByID(taskID uuid.UUID) (*task.Task, errors.Error)
	UpdateTask(_task task.Task) errors.Error
	DeleteTask(taskID uuid.UUID) errors.Error
}

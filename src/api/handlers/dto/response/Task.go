package response

import (
	"task_manager/src/core/domain/task"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func NewTask(_task task.Task) *Task {
	taskResponse := &Task{
		ID:          _task.ID(),
		Title:       _task.Title(),
		Description: _task.Description(),
		Status:      _task.Status(),
		CreatedAt:   _task.CreatedAt(),
		UpdatedAt:   _task.UpdatedAt(),
	}

	return taskResponse
}

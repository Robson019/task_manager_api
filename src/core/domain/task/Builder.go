package task

import (
	"github.com/google/uuid"
	"strings"
	"task_manager/src/core/errors"
	"task_manager/src/core/messages"
	"task_manager/src/utils/validator"
	"time"
)

type builder struct {
	task          *Task
	invalidFields []errors.InvalidField
}

func NewBuilder() *builder {
	return &builder{task: &Task{}}
}

func (instance *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidField{
			Name:        messages.TaskID,
			Description: messages.TaskIDErrorMessage,
		})
		return instance
	}
	instance.task.id = id
	return instance
}

func (instance *builder) WithTitle(title string) *builder {
	title = strings.TrimSpace(title)
	quantityOfWords := strings.Split(title, " ")
	if len(quantityOfWords) < 1 || len(title) < 2 {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidField{
			Name:        messages.TaskTitle,
			Description: messages.TaskTitleErrorMessage,
		})
		return instance
	}
	instance.task.title = title
	return instance
}

func (instance *builder) WithDescription(description string) *builder {
	description = strings.TrimSpace(description)
	instance.task.description = description

	return instance
}

func (instance *builder) WithStatus(status string) *builder {
	status = strings.TrimSpace(status)
	if status != messages.StatusPending && status != messages.StatusProgress &&
		status != messages.StatusDone {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidField{
			Name:        messages.TaskStatus,
			Description: messages.TaskStatusErrorMessage,
		})
		return instance
	}
	instance.task.status = status
	return instance
}

func (instance *builder) WithCreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidField{
			Name:        messages.TaskCreatedAt,
			Description: messages.TaskCreatedAtErrorMessage,
		})
		return instance
	}
	instance.task.createdAt = &createdAt
	return instance
}

func (instance *builder) WithUpdatedAt(updatedAt time.Time) *builder {
	instance.task.updatedAt = &updatedAt
	return instance
}

func (instance *builder) Build() (*Task, errors.Error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.NewValidationError(messages.TaskErrorMessage, instance.invalidFields...)
	}
	return instance.task, nil
}

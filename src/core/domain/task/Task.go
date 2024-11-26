package task

import (
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Task struct {
	id          uuid.UUID
	title       string
	description string
	status      string
	createdAt   *time.Time
	updatedAt   *time.Time
}

func (instance *Task) ID() uuid.UUID {
	return instance.id
}

func (instance *Task) Title() string {
	return instance.title
}

func (instance *Task) Description() string {
	return instance.description
}

func (instance *Task) Status() string {
	return instance.status
}

func (instance *Task) CreatedAt() *time.Time {
	return instance.createdAt
}

func (instance *Task) UpdatedAt() *time.Time {
	return instance.updatedAt
}

func (instance *Task) IsZero() bool {
	return reflect.DeepEqual(instance, &Task{})
}

package services

import (
	"github.com/google/uuid"
	"task_manager/src/core/domain/task"
	"task_manager/src/core/errors"
	"task_manager/src/core/errors/logger"
	"task_manager/src/core/interfaces/primary"
	"task_manager/src/core/interfaces/repository"
)

var _ primary.TaskManager = (*TaskServices)(nil)

type TaskServices struct {
	taskRepository repository.TaskLoader
	logger         logger.Logger
}

func (instance TaskServices) CreateTask(_task task.Task) (*uuid.UUID, errors.Error) {
	id, err := instance.taskRepository.CreateTask(_task)
	if err != nil {
		instance.logger.Log(err)
		return nil, err
	}
	return id, nil
}

func (instance TaskServices) FindTasks() ([]task.Task, errors.Error) {
	tasks, err := instance.taskRepository.FindTasks()
	if err != nil {
		instance.logger.Log(err)
		return nil, err
	}
	return tasks, nil
}

func (instance TaskServices) FindTaskByID(taskID uuid.UUID) (*task.Task, errors.Error) {
	_task, err := instance.taskRepository.FindTaskByID(taskID)
	if err != nil {
		instance.logger.Log(err)
		return nil, err
	}
	return _task, nil
}

func (instance TaskServices) UpdateTask(_task task.Task) errors.Error {
	_, err := instance.taskRepository.FindTaskByID(_task.ID())
	if err != nil {
		instance.logger.Log(err)
		return err
	}
	err = instance.taskRepository.UpdateTask(_task)
	if err != nil {
		instance.logger.Log(err)
		return err
	}
	return nil
}

func (instance TaskServices) DeleteTask(taskID uuid.UUID) errors.Error {
	_, err := instance.taskRepository.FindTaskByID(taskID)
	if err != nil {
		instance.logger.Log(err)
		return err
	}
	err = instance.taskRepository.DeleteTask(taskID)
	if err != nil {
		instance.logger.Log(err)
		return err
	}
	return nil
}

func NewTaskServices(taskRepository repository.TaskLoader, logger logger.Logger) *TaskServices {
	return &TaskServices{
		taskRepository: taskRepository,
		logger:         logger,
	}
}

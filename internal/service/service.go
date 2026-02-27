package service

import (
	"task-api/internal/model"
	"task-api/internal/storage"
)

type TaskService struct {
	storage *storage.MemoryStorage
}

func NewTaskService(s *storage.MemoryStorage) *TaskService {
	return &TaskService{storage: s}
}

func (s *TaskService) GetTasks() []model.Task {
	return s.storage.GetTasks()
}

func (s *TaskService) CreateTask(title string) model.Task {
	return s.storage.CreateTask(title)
}

func (s *TaskService) CompleteTask(id int) (model.Task, bool) {
	return s.storage.CompleteTask(id)
}

func (s *TaskService) DeleteTask(id int) bool {
	return s.storage.DeleteTask(id)
}

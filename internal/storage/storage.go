package storage

import (
	"sync"
	"task-api/internal/model"
)

type MemoryStorage struct {
	mu    sync.Mutex
	tasks map[int]model.Task
	id    int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks: make(map[int]model.Task),
		id:    1,
	}
}

func (s *MemoryStorage) GetTasks() []model.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks := make([]model.Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}
	return tasks
}

func (s *MemoryStorage) CreateTask(title string) model.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := model.Task{
		ID:        s.id,
		Title:     title,
		Completed: false,
	}

	s.tasks[s.id] = task
	s.id++

	return task
}

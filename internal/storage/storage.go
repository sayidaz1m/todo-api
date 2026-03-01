package storage

import (
	"sync"
	"task-api/internal/model"
	"time"
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
		CreatedAt: time.Now().UTC(),
	}

	s.tasks[s.id] = task
	s.id++

	return task
}

func (s *MemoryStorage) CompleteTask(id int) (model.Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, ok := s.tasks[id]
	if !ok {
		return model.Task{}, false
	}

	now := time.Now()
	task.Completed = true
	task.CompletedAt = &now
	s.tasks[id] = task
	return task, true
}

func (s *MemoryStorage) DeleteTask(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.tasks[id]; !ok {
		return false
	}
	delete(s.tasks, id)
	return true
}

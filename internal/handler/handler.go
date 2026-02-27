package handler

import (
	"encoding/json"
	"net/http"
	"task-api/internal/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.service.GetTasks()
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title string `json:"title"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	task := h.service.CreateTask(req.Title)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

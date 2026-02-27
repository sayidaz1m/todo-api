package app

import (
	"net/http"
	"task-api/internal/handler"
	"task-api/internal/service"
	"task-api/internal/storage"
)

type App struct {
	mux *http.ServeMux
}

func NewApp() *App {
	st := storage.NewMemoryStorage()
	svc := service.NewTaskService(st)
	h := handler.NewTaskHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetTasks(w, r)
		case http.MethodPost:
			h.CreateTask(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return &App{mux: mux}
}

func (a *App) Run() error {
	return http.ListenAndServe(":8080", a.mux)
}

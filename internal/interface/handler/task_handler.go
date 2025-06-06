package handler

import (
    "encoding/json"
    "errors"
    "io/ioutil"
    "net/http"

    "demo_go_clean_architecutre/internal/entity"
    "demo_go_clean_architecutre/internal/infrastructure/repository"
    "demo_go_clean_architecutre/internal/usecase"
)

type TaskHandler struct {
    usecase usecase.TaskUseCaseInterface
}

func NewTaskHandler(u usecase.TaskUseCaseInterface) *TaskHandler {
    return &TaskHandler{usecase: u}
}

func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        h.handleGetTasks(w, r)
    case "POST":
        h.handleCreateTask(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func (h *TaskHandler) handleGetTasks(w http.ResponseWriter, r *http.Request) {
    tasks, err := h.usecase.GetTasks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    jsonData, err := json.Marshal(tasks)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonData)
}

func (h *TaskHandler) handleCreateTask(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    var task entity.Task
    err = json.Unmarshal(body, &task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    err = h.usecase.CreateTask(task)
    if err != nil {
        if errors.Is(err, repository.ErrEmptyTitle) {
            http.Error(w, err.Error(), http.StatusBadRequest)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }
    w.WriteHeader(http.StatusCreated)
}

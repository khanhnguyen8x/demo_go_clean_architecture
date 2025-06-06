package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"demo_go_clean_architecutre/internal/entity"
	"demo_go_clean_architecutre/internal/interface/handler"
)

// mockTaskUseCase implements the minimal interface required by TaskHandler.
type mockTaskUseCase struct {
	tasks     []entity.Task
	createErr error
}

func (m *mockTaskUseCase) GetTasks() ([]entity.Task, error) {
	return m.tasks, nil
}

func (m *mockTaskUseCase) CreateTask(task entity.Task) error {
	return m.createErr
}

func TestHandleGetTasks(t *testing.T) {
	mockUC := &mockTaskUseCase{
		tasks: []entity.Task{{ID: 1, Title: "Task1", Done: false}},
	}
	taskHandler := handler.NewTaskHandler(mockUC)
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	taskHandler.HandleTasks(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	bodyBytes, _ := ioutil.ReadAll(rr.Body)
	var tasks []entity.Task
	if err := json.Unmarshal(bodyBytes, &tasks); err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
	if len(tasks) == 0 {
		t.Errorf("expected tasks in response")
	}
}

func TestHandleCreateTask_Success(t *testing.T) {
	mockUC := &mockTaskUseCase{}
	taskHandler := handler.NewTaskHandler(mockUC)
	newTask := entity.Task{Title: "New Task"}
	taskJSON, _ := json.Marshal(newTask)
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	taskHandler.HandleTasks(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusCreated)
	}
}

func TestHandleCreateTask_Error(t *testing.T) {
	simulatedErr := errors.New("simulated error")
	mockUC := &mockTaskUseCase{createErr: simulatedErr}
	taskHandler := handler.NewTaskHandler(mockUC)
	newTask := entity.Task{Title: ""}
	taskJSON, _ := json.Marshal(newTask)
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	taskHandler.HandleTasks(rr, req)

	if status := rr.Code; status != http.StatusBadRequest && status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v, want %v or %v", status, http.StatusBadRequest, http.StatusInternalServerError)
	}
}

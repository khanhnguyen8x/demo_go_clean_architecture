package usecase_test

import (
	"errors"
	"testing"

	"demo_go_clean_architecutre/internal/entity"
	"demo_go_clean_architecutre/internal/usecase"
)

// mockTaskRepository implements the TaskRepository interface for testing.
type mockTaskRepository struct {
	tasks     []entity.Task
	createErr error
}

func (m *mockTaskRepository) GetTasks() ([]entity.Task, error) {
	return m.tasks, nil
}

func (m *mockTaskRepository) CreateTask(task entity.Task) error {
	if m.createErr != nil {
		return m.createErr
	}
	task.ID = len(m.tasks) + 1
	m.tasks = append(m.tasks, task)
	return nil
}

func TestTaskUseCase_CreateTask_Success(t *testing.T) {
	mockRepo := &mockTaskRepository{}
	u := usecase.NewTaskUseCase(mockRepo)
	task := entity.Task{Title: "New Task"}
	err := u.CreateTask(task)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	tasks, _ := u.GetTasks()
	if len(tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(tasks))
	}
}

func TestTaskUseCase_CreateTask_Error(t *testing.T) {
	simulatedErr := errors.New("simulated repository error")
	mockRepo := &mockTaskRepository{createErr: simulatedErr}
	u := usecase.NewTaskUseCase(mockRepo)
	task := entity.Task{Title: "New Task"}
	err := u.CreateTask(task)
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
	if !errors.Is(err, simulatedErr) {
		t.Errorf("expected error to wrap simulated error, got: %v", err)
	}
}

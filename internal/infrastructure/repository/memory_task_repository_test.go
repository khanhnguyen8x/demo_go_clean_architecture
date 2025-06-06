package repository_test

import (
	"testing"

	"demo_go_clean_architecutre/internal/entity"
	"demo_go_clean_architecutre/internal/infrastructure/repository"
)

func TestMemoryTaskRepository_CreateTask(t *testing.T) {
	repo := repository.NewMemoryTaskRepository()

	// Test empty title returns error
	err := repo.CreateTask(entity.Task{Title: ""})
	if err == nil {
		t.Errorf("expected error for empty title, got nil")
	}
	if err.Error() != "task title cannot be empty" {
		t.Errorf("unexpected error message: %v", err)
	}

	// Test valid task creation
	task := entity.Task{Title: "Test Task"}
	err = repo.CreateTask(task)
	if err != nil {
		t.Errorf("expected nil error, got: %v", err)
	}

	tasks, _ := repo.GetTasks()
	if len(tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(tasks))
	}
	if tasks[0].ID != 1 {
		t.Errorf("expected task ID to be 1, got %d", tasks[0].ID)
	}
}

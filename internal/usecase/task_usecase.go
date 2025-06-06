package usecase

import (
    "demo_go_clean_architecutre/internal/entity"
    "fmt"
)

type TaskRepository interface {
    GetTasks() ([]entity.Task, error)
    CreateTask(task entity.Task) error
}

type TaskUseCase struct {
    repo TaskRepository
}

func NewTaskUseCase(repo TaskRepository) *TaskUseCase {
    return &TaskUseCase{repo: repo}
}

func (u *TaskUseCase) GetTasks() ([]entity.Task, error) {
    return u.repo.GetTasks()
}

func (u *TaskUseCase) CreateTask(task entity.Task) error {
    if err := u.repo.CreateTask(task); err != nil {
        return fmt.Errorf("failed to create task: %w", err)
    }
    return nil
}

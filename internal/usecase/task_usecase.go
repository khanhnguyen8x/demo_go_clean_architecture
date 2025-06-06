package usecase

import "demo_go_clean_architecutre/internal/entity"

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
    return u.repo.CreateTask(task)
}

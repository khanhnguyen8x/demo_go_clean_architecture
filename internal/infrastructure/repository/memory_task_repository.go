package repository

import (
    "errors"
    "demo_go_clean_architecutre/internal/entity"
)

var ErrEmptyTitle = errors.New("task title cannot be empty")

type MemoryTaskRepository struct {
    tasks  []entity.Task
    nextID int
}

func NewMemoryTaskRepository() *MemoryTaskRepository {
    return &MemoryTaskRepository{
        tasks:  make([]entity.Task, 0),
        nextID: 1,
    }
}

func (r *MemoryTaskRepository) GetTasks() ([]entity.Task, error) {
    return r.tasks, nil
}

func (r *MemoryTaskRepository) CreateTask(task entity.Task) error {
    if task.Title == "" {
        return ErrEmptyTitle
    }
    task.ID = r.nextID
    r.nextID++
    r.tasks = append(r.tasks, task)
    return nil
}

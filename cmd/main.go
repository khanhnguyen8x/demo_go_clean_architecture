package main

import (
    "log"
    "net/http"

    "demo_go_clean_architecutre/internal/infrastructure/repository"
    "demo_go_clean_architecutre/internal/interface/handler"
    "demo_go_clean_architecutre/internal/usecase"
)

func main() {
    repo := repository.NewMemoryTaskRepository()
    usecase := usecase.NewTaskUseCase(repo)
    taskHandler := handler.NewTaskHandler(usecase)

    http.HandleFunc("/tasks", taskHandler.HandleTasks)

    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}

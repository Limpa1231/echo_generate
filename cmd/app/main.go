package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"firstRest/internal/database"
	"firstRest/internal/handlers"
	"firstRest/internal/taskService"
	"firstRest/internal/web/tasks"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Создаем репозиторий, сервис и обработчики
	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	// Инициализация Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())  // Логирование запросов
	e.Use(middleware.Recover()) // Восстановление после паник

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(handler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)


	e.POST("/api/tasks", handler.AddTaskHandler)
	e.GET("/api/tasks", handler.ShowTasksHandler)
	e.PUT("/api/tasks/:id", handler.UpdateTaskHandler)
	e.DELETE("/api/tasks/:id", handler.DeleteTaskHandler)
	
	// Запуск сервера
	log.Println("Server started at localhost:8080")
	if err := e.Start("localhost:8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
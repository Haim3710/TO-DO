package main

import (
    "log"
    _"os"

    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "TO-DO/internal/database"
    "TO-DO/internal/handlers"
    "TO-DO/internal/repository"
)

func main() {
	// Загрузка переменных окружения из .env файла
    if err := godotenv.Load(); err != nil {
        log.Printf("Error loading .env file: %v", err)
        log.Fatal("Make sure the .env file exists and is accessible")
    }
    // Инициализация базы данных
    if err := database.InitDB(); err != nil {
        log.Fatal(err)
    }
    // Создание Fiber приложения
    app := fiber.New()

    taskRepo := repository.NewTaskRepository(database.DB)
    taskHandler := handlers.NewTaskHandler(taskRepo)

	// Регистрация маршрутов
    app.Post("/tasks", taskHandler.CreateTask)
    app.Get("/tasks", taskHandler.GetTasks)
    app.Put("/tasks/:id", taskHandler.UpdateTask)
    app.Delete("/tasks/:id", taskHandler.DeleteTask)

	// Запуск сервера
    log.Fatal(app.Listen(":3000"))
}
package handlers

import (
    "github.com/gofiber/fiber/v2"
    "TO-DO/internal/models"
    "TO-DO/internal/repository"
)

type TaskHandler struct {
    repo *repository.TaskRepository
}

func NewTaskHandler(repo *repository.TaskRepository) *TaskHandler {
    return &TaskHandler{repo: repo}
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
    var task models.Task
    if err := c.BodyParser(&task); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    if err := h.repo.CreateTask(&task); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
    tasks, err := h.repo.GetTasks()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(tasks)
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }

    var task models.Task
    if err := c.BodyParser(&task); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    task.ID = id
    if err := h.repo.UpdateTask(&task); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(task)
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }

    if err := h.repo.DeleteTask(id); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.SendStatus(fiber.StatusNoContent)
}
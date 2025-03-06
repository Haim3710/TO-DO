package repository

import (
    "context"
    "TO-DO/internal/models"

    "github.com/jackc/pgx/v4/pgxpool"
)

type TaskRepository struct {
    db *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) *TaskRepository {
    return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
    query := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
    return r.db.QueryRow(context.Background(), query, task.Title, task.Description, task.Status).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
}

func (r *TaskRepository) GetTasks() ([]models.Task, error) {
    query := `SELECT id, title, description, status, created_at, updated_at FROM tasks`
    rows, err := r.db.Query(context.Background(), query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var task models.Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }

    return tasks, nil
}

func (r *TaskRepository) UpdateTask(task *models.Task) error {
    query := `UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = now() WHERE id = $4`
    _, err := r.db.Exec(context.Background(), query, task.Title, task.Description, task.Status, task.ID)
    return err
}

func (r *TaskRepository) DeleteTask(id int) error {
    query := `DELETE FROM tasks WHERE id = $1`
    _, err := r.db.Exec(context.Background(), query, id)
    return err
}
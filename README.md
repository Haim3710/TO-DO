# TO-DO

Это REST API для управления задачами (TODO-лист), разработанное на Go с использованием Fiber и PostgreSQL.

## Запуск проекта

1. Установлен Go и PostgreSQL.
2. Создана база данных `todo_db` в PostgreSQL.
3. Применил миграции:

```bash
psql -U your_db_user -d todo_db -a -f migrations/001_create_tasks_table.sql

4. Создал файл .env и заполните его данными для подключения к базе данных.
5. Запустите сервер: ```go run cmd/main.go```
6. API Endpoints
__ POST /tasks – Создание задачи.

__ GET /tasks – Получение списка всех задач.

__ PUT /tasks/:id – Обновление задачи.

__ DELETE /tasks/:id – Удаление задачи.
```

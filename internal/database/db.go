package database

import (
    "context"
    "fmt"
    "os"

    "github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() error {
    var err error
    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    DB, err = pgxpool.Connect(context.Background(), connStr)
    if err != nil {
        return err
    }

    return nil
}
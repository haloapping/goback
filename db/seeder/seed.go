package main

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"os"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/goback/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
	pool := db.NewConnection(connString)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	err = GenerateFakeData(context.Background(), pool, 500, 500)
	if err != nil {
		panic(err)
	}
}

func GenerateFakeData(ctx context.Context, pool *pgxpool.Pool, nUser int, nTask int) error {
	userTx, err := pool.Begin(ctx)
	if err != nil {
		userTx.Rollback(ctx)

		return err
	}

	userIdxs := make([]string, 0)
	for range nUser {
		q := `
			INSERT INTO users(id, username, password)
			VALUES($1, $2, $3)
			RETURNING id;
		`
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(gofakeit.Password(true, true, true, true, false, 8)), bcrypt.DefaultCost)
		if err != nil {
			userTx.Rollback(ctx)

			return err
		}
		row := userTx.QueryRow(ctx, q, ulid.Make().String(), gofakeit.Username(), string(hashPassword))
		var id string
		err = row.Scan(&id)
		if err != nil {
			userTx.Rollback(ctx)

			return err
		}
		userIdxs = append(userIdxs, id)
	}
	userTx.Commit(ctx)

	taskTx, err := pool.Begin(ctx)
	if err != nil {
		taskTx.Rollback(ctx)

		return err
	}
	for range nTask {
		q := `
			INSERT INTO tasks(id, user_id, title, description)
			VALUES($1, $2, $3, $4)
			RETURNING id;
		`
		randomUserIdx := rand.IntN(len(userIdxs))
		row := taskTx.QueryRow(ctx, q, ulid.Make().String(), userIdxs[randomUserIdx], gofakeit.Sentence(3), gofakeit.Sentence(15))
		var id string
		err := row.Scan(&id)
		if err != nil {
			taskTx.Rollback(ctx)

			return err
		}
	}
	taskTx.Commit(ctx)

	return nil
}

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"os"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/goback/db"
	"github.com/jackc/pgx/v5"
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

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
	pool := db.NewConnection(connStr)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	err = GenerateFakeData(context.Background(), pool, 100, 100)
	if err != nil {
		panic(err)
	}
}

func GenerateFakeData(ctx context.Context, pool *pgxpool.Pool, nUser int, nTask int) error {
	// Batch insert users
	userBatch := &pgx.Batch{}
	idx := 0
	userIdxs := make([]string, nUser)
	for range nUser {
		id := ulid.Make().String()
		username := gofakeit.Username()
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(gofakeit.Password(true, true, true, true, false, 8)), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		userBatch.Queue(
			`INSERT INTO users(id, username, password) VALUES($1, $2, $3) RETURNING id;`,
			id, username, string(hashPassword),
		)
		userIdxs[idx] = id
		idx++
	}

	userResults := pool.SendBatch(ctx, userBatch)
	for range nUser {
		var id string
		if err := userResults.QueryRow().Scan(&id); err != nil {
			userResults.Close()

			return err
		}
	}
	if err := userResults.Close(); err != nil {
		return err
	}

	// Batch insert tasks
	taskBatch := &pgx.Batch{}
	for range nTask {
		id := ulid.Make().String()
		randomUserIdx := rand.IntN(len(userIdxs))
		taskBatch.Queue(
			`INSERT INTO tasks(id, user_id, title, description) VALUES($1, $2, $3, $4) RETURNING id;`,
			id, userIdxs[randomUserIdx], gofakeit.Sentence(3), gofakeit.Sentence(15),
		)
	}

	taskResults := pool.SendBatch(ctx, taskBatch)
	for range nTask {
		var id string
		if err := taskResults.QueryRow().Scan(&id); err != nil {
			taskResults.Close()
			return err
		}
	}
	if err := taskResults.Close(); err != nil {
		return err
	}

	return nil
}

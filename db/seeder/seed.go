package main

import (
	"context"
	"flag"
	"log"
	"math/rand/v2"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/goback/config"
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

	connStr := config.DBConnStr(".env")
	pool := db.NewConnection(connStr)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	nUser := flag.Int("nuser", 10, "number of fake user")
	nTask := flag.Int("ntask", 10, "number of fake task")
	flag.Parse()

	err = GenerateFakeData(context.Background(), pool, *nUser, *nTask)
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
	_, err := userResults.Exec()
	if err != nil {
		return err
	}
	err = userResults.Close()
	if err != nil {
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
	_, err = taskResults.Exec()
	if err != nil {
		return err
	}
	err = taskResults.Close()
	if err != nil {
		return err
	}

	return nil
}

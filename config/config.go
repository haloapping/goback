package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DBConnStr(envName string) string {
	_ = fmt.Sprintf("../%s", envName)
	err := godotenv.Load(fmt.Sprintf("./%s", envName))
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

	return connStr
}

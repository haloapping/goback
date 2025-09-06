package main

import (
	"fmt"
	"io"
	stdlog "log"
	"net/http"
	"os"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/goback/api/task"
	"github.com/goback/api/user"
	"github.com/goback/db"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

//	@title			Goback API
//	@version		1.0
//	@description	Goback API
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Alfiyanto Kondolele
//	@contact.url	https://haloapping.com
//	@contact.email	haloapping@gmail.com

// @BasePath	/
func main() {
	// Open file for logging
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		zlog.Fatal().Err(err).Msg("Failed to open log file")
	}
	defer logFile.Close()

	// Console writer
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05"}

	// MultiWriter -> write to both console and file
	multi := io.MultiWriter(consoleWriter, logFile)

	// Set global logger
	zlog.Logger = zerolog.New(multi).With().Timestamp().Logger()

	r := echo.New()
	r.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod: true,
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			sqlQuery, _ := c.Get("query").(string)
			fmt.Printf(
				"Method: %v, Uri: %v, Status: %v, Query: %v\n",
				v.Method, v.URI, v.Status, sqlQuery,
			)
			return nil
		},
	}))

	err = godotenv.Load(".env")
	if err != nil {
		stdlog.Fatal("Error loading .env file")
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

	userRepo := user.NewRepository(pool)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	user.Router(r.Group("/users"), userHandler)

	taskRepo := task.NewRepository(pool)
	taskService := task.NewService(taskRepo)
	taskHandler := task.NewHandler(taskService)
	task.Router(r.Group("/tasks"), taskHandler)

	r.GET("/", func(c echo.Context) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Goback API",
			},
			DarkMode:   true,
			Theme:      scalar.ThemeDeepSpace,
			HideModels: true,
		})

		if err != nil {
			fmt.Printf("%v", err)
		}

		return c.HTML(http.StatusOK, htmlContent)
	})

	r.Start(":3000")
}

package main

import (
	"fmt"
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/goback/api/task"
	"github.com/goback/api/user"
	"github.com/goback/config"
	"github.com/goback/db"
	customMiddleware "github.com/goback/middleware"
	"github.com/labstack/echo/v4"
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
	// console and file log
	logFile, err := customMiddleware.MultiLog()
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	// initiate database pooling
	connStr := config.DBConnStr(".env")
	pool := db.NewConnection(connStr)
	defer pool.Close()

	// routing
	r := echo.New()
	customMiddleware.EchoLogger(r)

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

package middleware

import (
	"fmt"
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func EchoLogger(r *echo.Echo) {
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
}

func MultiLog() (*os.File, error) {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05"}
	multi := io.MultiWriter(consoleWriter, logFile)
	zlog.Logger = zerolog.New(multi).With().Timestamp().Logger()

	return logFile, nil
}

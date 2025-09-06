package user

import "github.com/labstack/echo/v4"

func Router(g *echo.Group, h Handler) {
	g.POST("/register", h.Register)
	g.POST("/login", h.Login)
	g.GET("/biodata/:id", h.Biodata)
}

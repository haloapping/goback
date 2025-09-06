package task

import "github.com/labstack/echo/v4"

func Router(g *echo.Group, h Handler) {
	g.GET("", h.FindAll)
	g.GET("/:id", h.FindById)
	g.GET("/:id", h.FindByUserId)
	g.POST("", h.Add)
	g.PATCH("/:id", h.UpdateById)
	g.DELETE("/:id", h.DeleteById)
}

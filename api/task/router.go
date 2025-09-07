package task

import "github.com/labstack/echo/v4"

func Router(g *echo.Group, h Handler) {
	g.GET("", h.GetAll)
	g.GET("/:id", h.GetById)
	g.GET("/:userId", h.GetAllByUserId)
	g.POST("", h.Add)
	g.PATCH("/:id", h.UpdateById)
	g.DELETE("/:id", h.DeleteById)
}

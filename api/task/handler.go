package task

import (
	"net/http"
	"strconv"

	"github.com/goback/api"
	"github.com/labstack/echo/v4"
	zlog "github.com/rs/zerolog/log"
)

type Handler struct {
	Service
}

func NewHandler(s Service) Handler {
	return Handler{
		Service: s,
	}
}

// Add new task
//
//	@Summary		Add new task
//	@Description	Add new task
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			task		body		AddReq	true	"Add request"
//	@Success		200			{object}	api.SingleDataResp[Task]
//	@Router			/tasks   	[post]
func (h *Handler) Add(c echo.Context) error {
	var reqBody AddReq
	err := c.Bind(&reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return c.JSON(
			http.StatusBadRequest,
			api.ErrorResp{
				Error: err.Error(),
			},
		)
	}

	validationErr := AddValidation(reqBody)
	if len(validationErr) > 0 {
		zlog.Info().Interface("validation", validationErr).Msg("validation")

		return c.JSON(
			http.StatusBadRequest,
			api.ValidationResp{
				Validation: validationErr,
			},
		)
	}

	data, err := h.Service.Add(c, reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return c.JSON(
			http.StatusBadRequest,
			api.ErrorResp{
				Error: err.Error(),
			},
		)
	}

	zlog.Info().Msg("task is created")

	return c.JSON(
		http.StatusCreated,
		api.SingleDataResp[Task]{
			Message: "Task is created",
			Data:    data,
		},
	)
}

// Find task by id
//
//	@Summary		Find task by id
//	@Description	Find task by id
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"task id"
//	@Success		200				{object}	api.SingleDataResp[Task]
//	@Router			/tasks/{id} 	[get]
func (h *Handler) FindById(c echo.Context) error {
	id := c.Param("id")
	if id == "{id}" {
		validation := map[string][]string{
			"id": {"cannot empty"},
		}

		zlog.Info().Interface("validation", validation).Msg("validation")

		return c.JSON(
			http.StatusBadRequest,
			api.ValidationResp{
				Validation: validation,
			},
		)
	}

	data, err := h.Service.FindById(c, id)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return c.JSON(
			http.StatusBadRequest,
			api.ErrorResp{
				Error: err.Error(),
			},
		)
	}

	zlog.Info().Msg("success to retrieve task by id")

	return c.JSON(
		http.StatusCreated,
		api.SingleDataResp[Task]{
			Message: "Success to retrieve task by id",
			Data:    data,
		},
	)
}

// Find all tasks by user id
//
//	@Summary		Find all tasks by user id
//	@Description	Find all tasks by user id
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"user id"
//	@Param			offset			query		int		true	"offset"	default(1)
//	@Param			limit			query		int		true	"limit"		default(15)
//	@Success		200				{object}	api.MultipleDataResp[Task]
//	@Router			/tasks/{id} 	[get]
func (h *Handler) FindAllByUserId(c echo.Context) error {
	validationErr := make(map[string][]string)

	id := c.Param("id")
	if id == "{id}" {
		validationErr["id"] = append(validationErr["id"], "cannot empty")
	}

	offsetStr := c.QueryParam("offset")
	if offsetStr == "" {
		validationErr["offset"] = append(validationErr["offset"], "cannot empty")
	}

	limitStr := c.QueryParam("limit")
	if limitStr == "" {
		validationErr["limit"] = append(validationErr["limit"], "cannot empty")
	}

	if len(validationErr) > 0 {
		zlog.Info().Interface("validation", validationErr).Msg("validation")

		return c.JSON(
			http.StatusBadRequest,
			api.ValidationResp{
				Validation: validationErr,
			},
		)
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			api.ErrorResp{
				Error: err.Error(),
			},
		)
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			api.ErrorResp{
				Error: err.Error(),
			},
		)
	}

	data, err := h.Service.FindAllByUserId(c, id, limit, offset)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return c.JSON(
			http.StatusBadRequest,
			api.ErrorResp{
				Error: err.Error(),
			},
		)
	}

	zlog.Info().Msg("success to retrieve task by user id")

	return c.JSON(
		http.StatusCreated,
		api.MultipleDataResp[UserTask]{
			Message:   "Success to retrieve task by user id",
			TotalItem: limit,
			Data:      data,
		},
	)
}

// Find all tasks
//
//	@Summary		Find all tasks
//	@Description	Find all tasks
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			offset		query		int	true	"offset"	default(1)
//	@Param			limit		query		int	true	"limit"		default(15)
//	@Success		200			{object}	api.MultipleDataResp[Task]
//	@Router			/tasks   	[get]
func (h *Handler) FindAll(c echo.Context) error {
	validationErr := make(map[string][]string)

	offsetStr := c.QueryParam("offset")
	if offsetStr == "" {
		validationErr["offset"] = append(validationErr["offset"], "cannot empty")
	}

	limitStr := c.QueryParam("limit")
	if limitStr == "" {
		validationErr["limit"] = append(validationErr["limit"], "cannot empty")
	}

	if len(validationErr) > 0 {
		zlog.Info().Interface("validation", validationErr).Msg("validation")

		return c.JSON(
			http.StatusBadRequest,
			api.ValidationResp{
				Validation: validationErr,
			},
		)
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			api.ErrorResp{
				Error: err.Error(),
			},
		)
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			api.ErrorResp{
				Error: err.Error(),
			},
		)
	}

	items, err := h.Service.FindAll(c, limit, offset)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return c.JSON(
			http.StatusBadRequest,
			api.ErrorResp{
				Error: err.Error(),
			},
		)
	}

	zlog.Info().Msg("success to retrieve all tasks")

	return c.JSON(
		http.StatusCreated,
		api.MultipleDataResp[Task]{
			Message:   "Success to retrieve all tasks",
			TotalItem: limit,
			Data:      items,
		},
	)
}

// Update task by id
//
//	@Summary		Update task by id
//	@Description	Update task by id
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string		true	"task id"
//	@Param			task			body		UpdateReq	true	"Update request"
//	@Success		200				{object}	api.SingleDataResp[Task]
//	@Router			/tasks/{id} 	[patch]
func (h *Handler) UpdateById(c echo.Context) error {
	id := c.Param("id")
	if id == "{id}" {
		validaion := map[string][]string{
			"id": {"cannot empty"},
		}

		zlog.Info().Interface("validation", validaion).Msg("validation")

		return c.JSON(
			http.StatusBadRequest,
			api.ValidationResp{
				Validation: validaion,
			},
		)
	}

	var reqBody UpdateReq
	err := c.Bind(&reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return c.JSON(
			http.StatusBadRequest,
			api.ErrorResp{
				Error: err.Error(),
			},
		)
	}

	data, err := h.Service.UpdateById(c, id, reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return c.JSON(
			http.StatusBadRequest,
			api.ErrorResp{
				Error: err.Error(),
			},
		)
	}

	zlog.Info().Msg("success to update task by id")

	return c.JSON(
		http.StatusCreated,
		api.SingleDataResp[Task]{
			Message: "Success to update task by id",
			Data:    data,
		},
	)
}

// Delete task by id
//
//	@Summary		Delete task by id
//	@Description	Delete task by id
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"task id"
//	@Success		200				{object}	api.SingleDataResp[Task]
//	@Router			/tasks/{id} 	[delete]
func (h *Handler) DeleteById(c echo.Context) error {
	id := c.Param("id")
	if id == "{id}" {
		validation := map[string][]string{
			"id": {"cannot empty"},
		}

		zlog.Info().Interface("validation", validation).Msg("validation")

		return c.JSON(
			http.StatusBadRequest,
			api.ValidationResp{
				Validation: validation,
			},
		)
	}

	data, err := h.Service.DeleteById(c, id)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return c.JSON(
			http.StatusBadRequest,
			api.ErrorResp{
				Error: err.Error(),
			},
		)
	}

	zlog.Info().Msg("success to delete task by id")

	return c.JSON(
		http.StatusCreated,
		api.SingleDataResp[Task]{
			Message: "Success to delete task by id",
			Data:    data,
		},
	)
}

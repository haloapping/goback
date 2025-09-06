package user

import (
	"net/http"

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

// Register user
//
//	@Summary		Register user
//	@Description	Register user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			register			user		body	UserRegisterReq	true	"Register user request"
//	@Success		200					{object}	api.SingleDataResp[UserRegister]
//	@Router			/users/register   	[post]
func (h *Handler) Register(c echo.Context) error {
	var reqBody UserRegisterReq
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

	validationErr := RegisterValidation(reqBody)
	if len(validationErr) > 0 {
		zlog.Info().Fields(validationErr).Msg("validationErr")

		return c.JSON(
			http.StatusBadRequest,
			api.ValidationResp{
				Validation: validationErr,
			},
		)
	}

	u, err := h.Service.Register(c, reqBody)
	if err != nil {
		zlog.Error().Msg(err.Error())

		return c.JSON(
			http.StatusBadRequest,
			api.ErrorResp{
				Error: err.Error(),
			},
		)
	}

	zlog.Info().Msg("user is registered")

	return c.JSON(
		http.StatusCreated,
		api.SingleDataResp[UserRegister]{
			Message: "user is registered",
			Data:    u,
		},
	)
}

// Login user
//
//	@Summary		Login user
//	@Description	Login user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			login			user		body	UserLoginReq	true	"Login user request"
//	@Success		200				{object}	api.SingleDataResp[UserLogin]
//	@Router			/users/login   	[post]
func (h *Handler) Login(c echo.Context) error {
	var reqBody UserLoginReq
	err := c.Bind(&reqBody)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			api.ErrorResp{
				Error: err.Error(),
			},
		)
	}

	validationErr := LoginValidation(reqBody)
	if len(validationErr) > 0 {
		return c.JSON(
			http.StatusBadRequest,
			api.ValidationResp{
				Validation: validationErr,
			},
		)
	}

	token, err := h.Service.Login(c, reqBody)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			api.ErrorResp{
				Error: err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusCreated,
		api.SingleDataResp[string]{
			Message: "login successfuly",
			Data:    token,
		},
	)
}

// Find biodata of user by id
//
//	@Summary		Find biodata of user by id
//	@Description	Find biodata of user by id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id						path		string	true	"biodata of user id"
//	@Success		200						{object}	api.SingleDataResp[UserBiodata]
//	@Router			/users/biodata/{id} 	[get]
func (h *Handler) Biodata(c echo.Context) error {
	id := c.Param("id")
	if id == "{id}" {
		return c.JSON(
			http.StatusBadRequest,
			api.ValidationResp{
				Validation: map[string][]string{
					"id": {"cannot empty"},
				},
			},
		)
	}

	b, err := h.Service.Biodata(c, id)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			api.ErrorResp{
				Error: err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusCreated,
		api.SingleDataResp[UserBiodata]{
			Message: "retrieve user biodata",
			Data:    b,
		},
	)
}

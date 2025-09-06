package api

type ValidationResp struct {
	Validation map[string][]string `json:"validation" binding:"required" extensions:"x-order=1"`
}

type ErrorResp struct {
	Error string `json:"error" binding:"required" extensions:"x-order=1"`
}

type SingleDataResp[data any] struct {
	Message string `json:"message" binding:"required" extensions:"x-order=1"`
	Data    data   `json:"data" binding:"required" extensions:"x-order=2"`
}

type MultipleDataResp[data any] struct {
	Message   string `json:"message" binding:"required" extensions:"x-order=1"`
	TotalItem int    `json:"totalItem" binding:"required" extensions:"x-order=2"`
	Data      []data `json:"data" binding:"required" extensions:"x-order=3"`
}

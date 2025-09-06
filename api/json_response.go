package api

type ValidationResp struct {
	Validation map[string][]string `json:"validation"`
}

type ErrorResp struct {
	Error string `json:"error"`
}

type SingleDataResp[data any] struct {
	Message string `json:"message"`
	Data    data   `json:"data"`
}

type MultipleDataResp[data any] struct {
	Message string `json:"message"`
	Data    []data `json:"data"`
}

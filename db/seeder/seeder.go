package seeder

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/goback/api/task"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

type Service struct {
	task.Repository
}

func (s *Service) InitSeed(c echo.Context, count int) error {
	for i := 0; i < count; i++ {
		req := task.AddReq{
			UserId:      ulid.Make().String(),
			Title:       gofakeit.Sentence(2),
			Description: gofakeit.Sentence(15),
		}

		_, err := s.Add(c, req)
		if err != nil {
			return err
		}
	}

	return nil
}

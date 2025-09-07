package task

import (
	"github.com/labstack/echo/v4"
)

type Service struct {
	Repository
}

func NewService(r Repository) Service {
	return Service{
		Repository: r,
	}
}

func (s *Service) Add(c echo.Context, req AddReq) (Task, error) {
	item, err := s.Repository.Add(c, req)
	if err != nil {
		return Task{}, err
	}

	return item, nil
}

func (s *Service) GetById(c echo.Context, id string) (Task, error) {
	item, err := s.Repository.GetById(c, id)
	if err != nil {
		return Task{}, err
	}

	return item, nil
}

func (s *Service) GetAllByUserId(c echo.Context, id string, limit int, offset int) (ut []UserTask, err error) {
	item, err := s.Repository.GetAllByUserId(c, id, limit, offset)
	if err != nil {
		return []UserTask{}, err
	}

	return item, nil
}

func (s *Service) GetAll(c echo.Context, limit int, offset int) (t []Task, err error) {
	items, err := s.Repository.GetAll(c, limit, offset)
	if err != nil {
		return []Task{}, err
	}

	return items, nil
}

func (s *Service) UpdateById(c echo.Context, id string, req UpdateReq) (Task, error) {
	item, err := s.Repository.UpdateById(c, id, req)
	if err != nil {
		return Task{}, err
	}

	return item, nil
}

func (s *Service) DeleteById(c echo.Context, id string) (Task, error) {
	item, err := s.Repository.DeleteById(c, id)
	if err != nil {
		return Task{}, err
	}

	return item, nil
}

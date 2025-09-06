package task

import "github.com/labstack/echo/v4"

type Service struct {
	Repository
}

func NewService(r Repository) Service {
	return Service{
		Repository: r,
	}
}

func (s *Service) Add(c echo.Context, req AddReq) (Task, error) {
	res, err := s.Repository.Add(c, req)
	if err != nil {
		return Task{}, err
	}

	return res, nil
}

func (s *Service) FindById(c echo.Context, id string) (Task, error) {
	res, err := s.Repository.FindById(c, id)
	if err != nil {
		return Task{}, err
	}

	return res, nil
}

func (s *Service) FindByUserId(c echo.Context, id string) ([]UserTask, error) {
	res, err := s.Repository.FindAllTasksByUserId(c, id)
	if err != nil {
		return []UserTask{}, err
	}

	return res, nil
}

func (s *Service) FindAll(c echo.Context) ([]Task, error) {
	res, err := s.Repository.FindAll(c)
	if err != nil {
		return []Task{}, err
	}

	return res, nil
}

func (s *Service) UpdateById(c echo.Context, id string, req UpdateReq) (Task, error) {
	res, err := s.Repository.UpdateById(c, id, req)
	if err != nil {
		return Task{}, err
	}

	return res, nil
}

func (s *Service) DeleteById(c echo.Context, id string) (Task, error) {
	res, err := s.Repository.DeleteById(c, id)
	if err != nil {
		return Task{}, err
	}

	return res, nil
}

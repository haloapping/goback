package user

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repository
}

func NewService(r Repository) Service {
	return Service{
		Repository: r,
	}
}

func (s *Service) Register(c echo.Context, req UserRegisterReq) (UserRegister, error) {
	byteHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserRegister{}, err
	}
	req.Password = string(byteHash)

	u, err := s.Repository.Register(c, req)
	if err != nil {
		return UserRegister{}, err
	}

	return u, nil
}

func (s *Service) Login(c echo.Context, req UserLoginReq) (string, error) {
	token, err := s.Repository.Login(c, req)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) Biodata(c echo.Context, id string) (UserBiodata, error) {
	u, err := s.Repository.Biodata(c, id)
	if err != nil {
		return UserBiodata{}, err
	}

	return u, nil
}

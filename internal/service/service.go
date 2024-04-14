package service

import (
	"auth_service/internal/entity"
	"auth_service/internal/repository"
)

type Auth interface {
	GenerateToken(userID int) (string, error)
	ParseToken(accessToken string) (int, error)
}

type User interface {
	SendEmailCode(email string) error
	CheckEmailCode(email, code string) error
	VerifyEmailCode(email, code string) error
	CreateUser(input *entity.UserRegInput) (int, error)
}

type City interface {
	AddCity(city string) error
	Search(query string) ([]entity.City, error)
}

type Service struct {
	Auth
	User
	City
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
		City: NewCityService(repos.City),
		Auth: NewAuthService(repos.User),
	}
}

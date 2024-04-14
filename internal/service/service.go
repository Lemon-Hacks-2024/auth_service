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

type StoreAddresses interface {
	AddStoreAddress(address string) error
}

type Ticket interface {
	CheckPrice(input *entity.TicketInput) error
}

type Service struct {
	Auth
	User
	City
	StoreAddresses
	Ticket
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:           NewUserService(repos.User),
		City:           NewCityService(repos.City),
		StoreAddresses: NewStoreAddressesService(repos.StoreAddresses),
		Auth:           NewAuthService(repos.User),
		Ticket:         NewTicketService(repos.Ticket),
	}
}

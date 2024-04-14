package repository

import (
	"auth_service/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
}

type User interface {
	CreateUser(user *entity.UserRegInput) (int, error)
}

type City interface {
	AddCity(name string) error
	Search(query string) ([]entity.City, error)
}

type Repository struct {
	Auth
	User
	City
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserRepo(db),
		City: NewCityRepo(db),
	}
}

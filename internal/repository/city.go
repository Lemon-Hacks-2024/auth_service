package repository

import (
	"auth_service/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CityRepo struct {
	db *sqlx.DB
}

func NewCityRepo(db *sqlx.DB) *CityRepo {
	return &CityRepo{db: db}
}

func (r *CityRepo) AddCity(name string) error {

	query := `INSERT INTO cities (name) values ($1)`
	_, err := r.db.Exec(query, name)

	if err != nil {
		return err
	}

	return nil

}

func (r *CityRepo) Search(searchQuery string) ([]entity.City, error) {
	var cities []entity.City

	query := `SELECT id, name FROM cities WHERE lower(name) LIKE lower($1)`
	err := r.db.Select(&cities, query, "%"+searchQuery+"%")
	if err != nil {
		return nil, err
	}

	return cities, nil
}

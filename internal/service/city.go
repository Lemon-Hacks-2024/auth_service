package service

import (
	"auth_service/internal/entity"
	"auth_service/internal/repository"
)

type CityService struct {
	cityRepo repository.City
}

func NewCityService(cityRepo repository.City) *CityService {
	return &CityService{cityRepo: cityRepo}
}

func (s *CityService) AddCity(city string) error {
	return s.cityRepo.AddCity(city)
}

func (s *CityService) Search(query string) ([]entity.City, error) {
	return s.cityRepo.Search(query)
}

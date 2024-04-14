package service

import (
	"auth_service/internal/entity"
	"auth_service/internal/repository"
	"fmt"
)

type UserService struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(input *entity.UserRegInput) (int, error) {

	err := s.VerifyEmailCode(input.Email, input.EmailCode)
	if err != nil {
		return 0, err
	}

	if input.PhoneNumberCode != "0000" {
		return 0, fmt.Errorf("phone_number_code неверен")
	}

	return s.userRepo.CreateUser(input)
}

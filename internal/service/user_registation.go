package service

import (
	"auth_service/internal/utils"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type emailCodes struct {
	Code      string
	ExpiresAt time.Time
}

var userEmailCodes = make(map[string]emailCodes)

func (s *UserService) SendEmailCode(email string) error {
	// Генерация 4-значного кода
	code := rand.Intn(10000)
	log.Println(code)

	// Сохранение кода в базу данных
	userEmailCodes[email] = emailCodes{
		Code:      fmt.Sprintf("%04d", code),
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}

	err := utils.SendEmail(email, code)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) CheckEmailCode(email, code string) error {

	// Поиск кода в базе данных
	ec, ok := userEmailCodes[email]
	if !ok {
		return fmt.Errorf("код не найден")
	}

	// Проверка кода
	if ec.Code != code {
		return fmt.Errorf("код неверен")
	}

	// Проверка истечения срока действия кода
	if time.Now().After(ec.ExpiresAt) {
		return fmt.Errorf("код истек")
	}

	return nil
}

func (s *UserService) VerifyEmailCode(email, code string) error {

	// Поиск кода в базе данных
	ec, ok := userEmailCodes[email]
	if !ok {
		return fmt.Errorf("код не найден")
	}

	// Проверка кода
	if ec.Code != code {
		return fmt.Errorf("код неверен")
	}

	// Проверка истечения срока действия кода
	if time.Now().After(ec.ExpiresAt) {
		return fmt.Errorf("код истек")
	}

	// Удаление кода из базы данных
	delete(userEmailCodes, email)

	return nil
}

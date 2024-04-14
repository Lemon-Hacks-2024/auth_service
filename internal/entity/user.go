package entity

import (
	"fmt"
)

type User struct {
	ID           int      `json:"user_id,omitempty" db:"id_user"`
	Name         string   `json:"name,omitempty" db:"name"`
	Surname      string   `json:"surname,omitempty" db:"surname"`
	CityID       int      `json:"city_id,omitempty" db:"city_id"`
	Email        string   `json:"email,omitempty" db:"email"`
	PhoneNumber  string   `json:"phone_number,omitempty" db:"phone_number"`
	Password     string   `json:"password,omitempty" db:"password"`
	PasswordHash string   `json:"password_hash,omitempty" db:"password_hash"`
	ImeiID       []string `json:"imei_id,omitempty" db:"id_imei"`
	AccessToken  string   `json:"access_token,omitempty"`
}

type UserAuthInput struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserEmailCodeInput struct {
	Email     string `json:"email,omitempty"`
	EmailCode string `json:"email_code,omitempty"`
}

func (u *UserEmailCodeInput) Validate() error {
	if u.Email == "" {
		return fmt.Errorf("email не может быть пустым")
	}
	if u.EmailCode == "" {
		return fmt.Errorf("email_code не может быть пустым")
	}
	return nil
}

type UserRegInput struct {
	Name            string `json:"name,omitempty"`
	Surname         string `json:"surname,omitempty"`
	CityID          int    `json:"city_id,omitempty"`
	Email           string `json:"email,omitempty"`
	EmailCode       string `json:"email_code,omitempty"`
	PhoneNumber     string `json:"phone_number,omitempty"`
	PhoneNumberCode string `json:"phone_number_code,omitempty"`
	Password        string `json:"password,omitempty"`
	PasswordHash    string `json:"password_hash,omitempty" db:"password_hash"`
	ImeiID          string `json:"imei_id,omitempty"`
}

func (u *UserRegInput) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("name не может быть пустым")
	}
	if u.Surname == "" {
		return fmt.Errorf("surname не может быть пустым")
	}
	if u.CityID == 0 {
		return fmt.Errorf("city не может быть пустым и должен быть числом не равным нулю")
	}
	if u.Email == "" {
		return fmt.Errorf("email не может быть пустым")
	}
	if u.EmailCode == "" {
		return fmt.Errorf("email_code не может быть пустым")
	}

	phoneNumber := ""
	for i := 0; i < len(u.PhoneNumber); i++ {
		switch u.PhoneNumber[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			phoneNumber += string(u.PhoneNumber[i])
		}
	}
	u.PhoneNumber = phoneNumber
	if u.PhoneNumber == "" {
		return fmt.Errorf("phone_number не может быть пустым")
	}

	if u.PhoneNumberCode == "" {
		return fmt.Errorf("phone_number_code не может быть пустым")
	}
	if u.Password == "" {
		return fmt.Errorf("password не может быть пустым")
	}
	if u.ImeiID == "" {
		return fmt.Errorf("imei_id не может быть пустым")
	}

	return nil
}

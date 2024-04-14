package repository

import (
	"auth_service/internal/entity"
	"crypto/sha1"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Транзакция для создания пользователя в базе данных при регистрации и запись imei в базу данных
func (r *UserRepo) CreateUser(user *entity.UserRegInput) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var userID int

	err = tx.QueryRow("INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id_user", user.Email, hashPassword(user.Password)).Scan(&userID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec("INSERT INTO imeis (id_user, imei) VALUES ($1, $2)", userID, user.ImeiID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func hashPassword(password string) string {
	sha := sha1.New()
	sha.Write([]byte(password))
	shaHash := sha.Sum(nil)
	shaHashString := fmt.Sprintf("%x", shaHash)

	return shaHashString
}

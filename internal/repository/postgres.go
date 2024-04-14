package repository

import (
	"auth_service/configs"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func NewPostgresDB(cfg configs.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPass, cfg.DBSSLMode))
	if err != nil {
		return nil, err
	}
	//
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	// Создание пула соединений
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	//
	log.Println("Подключение к БД прошло успешно")
	return db, nil
}

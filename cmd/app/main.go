package main

import (
	"auth_service/configs"
	"auth_service/internal/handler"
	"auth_service/internal/repository"
	"auth_service/internal/service"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	cfg, err := configs.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cfg)
	//
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Ошибка при инициализации БД: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	handlers.InitRoutes(cfg.AppPort)
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
}

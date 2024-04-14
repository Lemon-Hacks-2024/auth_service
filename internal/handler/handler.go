package handler

import (
	"auth_service/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(port string) {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	//
	api := router.Group("/api")
	{

		//api.GET("/test", h.TestFunc)

		api.POST("/login", h.AuthUser)
		api.POST("/registration", h.RegistrationUser)

		verifyCode := api.Group("/verify-code")
		{
			verifyCode.GET("/email/:email", h.SendEmailCode)
			verifyCode.POST("/email", h.CheckEmailCode)
			verifyCode.GET("/phone/:phone", nil)
			verifyCode.POST("/phone", nil)
		}

		tickets := api.Group("/prices")
		{
			tickets.POST("/check", h.CheckPrice)
		}

		cities := api.Group("/cities")
		{
			cities.GET("/", h.SearchCities)
		}

	}

	fmt.Println("Start auth_service on  http://127.0.0.1:" + port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

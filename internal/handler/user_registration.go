package handler

import (
	"auth_service/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) RegistrationUser(ctx *gin.Context) {
	var input entity.UserRegInput
	// Получение тела запроса
	if err := ctx.BindJSON(&input); err != nil {
		resp := Response{
			Message: "Неверное тело запроса",
		}
		resp.SendError(ctx, err, 400)
		return
	}
	// Валидация
	if err := input.Validate(); err != nil {
		resp := Response{
			Message: err.Error(),
		}
		resp.Send(ctx, 400)
		return
	}
	// Создание пользователя
	userID, err := h.services.User.CreateUser(&input)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			resp := Response{
				Message: "Пользователь с таким email уже существует",
			}
			resp.Send(ctx, 409)
			return
		}
		resp := Response{
			Message: "Не удалось создать пользователя",
		}
		resp.SendError(ctx, err, 500)
		return
	}

	// Создание JWT токена
	token, err := h.services.Auth.GenerateToken(userID)
	if err != nil {
		resp := Response{
			Message: "Не удалось создать токен",
		}
		resp.SendError(ctx, err, 500)
		return
	}

	// Отправка ответа
	resp := Response{
		Message: "Пользователь успешно создан",
		Details: map[string]string{
			"access_token": "Bearer " + token,
		},
	}
	resp.Send(ctx, 201)
	return
}

func (h *Handler) SendEmailCode(ctx *gin.Context) {
	emailInput := ctx.Param("email")
	log.Println(emailInput)
	// Валидация
	if emailInput == "" {
		resp := Response{
			Message: "Поле email не может быть пустым",
		}
		resp.Send(ctx, 400)
		return
	}
	// Отправка кода подтверждения почты на почту пользователя
	err := h.services.User.SendEmailCode(emailInput)
	if err != nil {
		resp := Response{
			Message: "Не удалось отправить код подтверждения почты",
		}
		resp.SendError(ctx, err, 500)
		return
	}
	// Отправка ответа
	resp := Response{
		Message: "Код подтверждения отправлен",
	}
	resp.Send(ctx, 200)
}

func (h *Handler) CheckEmailCode(ctx *gin.Context) {
	//	Получение тела запроса
	var input entity.UserEmailCodeInput
	if err := ctx.BindJSON(&input); err != nil {
		resp := Response{
			Message: "Неверное тело запроса",
		}
		resp.SendError(ctx, err, 400)
		return
	}
	//	Валидация
	if err := input.Validate(); err != nil {
		resp := Response{
			Message: err.Error(),
		}
		resp.Send(ctx, 400)
		return
	}
	//	Проверка кода подтверждения
	err := h.services.User.CheckEmailCode(input.Email, input.EmailCode)
	if err != nil {
		resp := Response{
			Message: err.Error(),
		}
		resp.SendError(ctx, err, 500)
		return
	}
	//	Отправка ответа
	resp := Response{
		Message: "Код подтверждения верен",
	}
	resp.Send(ctx, 200)
	return
}

package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) SearchCities(ctx *gin.Context) {
	// получение query параметра
	query := ctx.Query("query")
	if len(query) <= 3 {
		resp := Response{
			Message: "Поисковая строка должна быть больше трех символов",
			Details: nil,
		}
		resp.Send(ctx, 400)
		return
	}
	// Вызов сервиса
	cities, err := h.services.City.Search(query)
	if err != nil {
		resp := Response{
			Message: "Произошла ошибка при поиске городов",
			Details: err.Error(),
		}
		resp.SendError(ctx, err, 500)
		return
	}
	// Отправка ответа
	resp := Response{
		Message: "OK",
		Details: cities,
	}
	resp.Send(ctx, 200)
	return
}

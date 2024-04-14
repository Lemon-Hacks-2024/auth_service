package handler

import (
	"auth_service/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckPrice(c *gin.Context) {
	var ticket entity.TicketInput
	// Получение тела запроса
	if err := c.BindJSON(&ticket); err != nil {
		resp := Response{
			Message: "Неверное тело запроса",
		}
		resp.SendError(c, err, 400)
		return
	}
	// Валидация
	if ticket.NameProduct == "" || ticket.PriceProduct == 0 {
		resp := Response{
			Message: "Неверное тело запроса",
		}
		resp.Send(c, 400)
		return
	}
	// Проверка цены
	err := h.services.Ticket.CheckPrice(&ticket)
	if err != nil {
		resp := Response{
			Message: err.Error(),
		}
		resp.Send(c, 400)
		return
	}

	// Отправка ответа
	resp := Response{
		Message: "Цена прошла проверку",
	}
	resp.Send(c, 200)
	return
}

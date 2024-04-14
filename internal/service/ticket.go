package service

import (
	"auth_service/internal/entity"
	"auth_service/internal/repository"
	"fmt"
	"log"
)

type TicketService struct {
	ticketRepo repository.Ticket
}

func NewTicketService(ticketRepo repository.Ticket) *TicketService {
	return &TicketService{ticketRepo: ticketRepo}
}

func (s *TicketService) CheckPrice(input *entity.TicketInput) error {

	maxSocialPrice, err := s.ticketRepo.GetCocialPriceProductByName(input.NameProduct)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Не удалось получить цену продукта")
	}

	if input.PriceProduct <= maxSocialPrice {
		return fmt.Errorf("Цена продукта не превышает предельную цену соц.продукта. Предельная цена соц.продукта: %f", maxSocialPrice)
	}

	if maxSocialPrice == 0 {
		return fmt.Errorf("Не удалось получить цену соц.продукта")
	}

	return nil
}

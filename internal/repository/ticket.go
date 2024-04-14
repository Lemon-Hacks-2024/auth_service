package repository

import (
	"auth_service/internal/entity"
	"github.com/jmoiron/sqlx"
)

type TicketRepo struct {
	db *sqlx.DB
}

func NewTicketRepo(db *sqlx.DB) *TicketRepo {
	return &TicketRepo{db: db}
}

// СОздаём ticket в бд
func (r *TicketRepo) CreateTicket(input *entity.TicketInput) error {
	// Создаём ticket в бд c транзакцией

	//tx, err := r.db.Begin()
	//if err != nil {
	//	return err
	//}
	//
	//var ticketID int
	//
	//query := `INSERT INTO tickets (status, "refusedDescription") values ($1, $2)`
	//row, err := r.db.Exec(query, "open", "")
	//if err != nil {
	//	return err
	//}

	//err = row.LastInsertId(&ticketID)
	//if err != nil {
	//	return err
	//}
	//
	//// Добавляем ticket_id в tickets_details
	//queryDetails := `INSERT INTO tickets_details
	//	(ticket_id, reportedUser, ticketCity, sellerName, storeAddress, itemName, recommendedPrice, currentPrice)
	//	values ($1, $2, $3, $4, $5, $6, $7, $8)`

	return nil
}

func (r *TicketRepo) GetCocialPriceProductByName(nameProduct string) (float64, error) {

	query := `SELECT current_price_social_product
FROM social_products_price
    JOIN public.social_products sp
    on sp.id_social_product = social_products_price.id_social_product
where lower(sp.name_product) LIKE lower($1)`

	var price float64

	err := r.db.Get(&price, query, "%"+nameProduct+"%")

	if err != nil {
		return 0, err
	}

	return price, nil
}

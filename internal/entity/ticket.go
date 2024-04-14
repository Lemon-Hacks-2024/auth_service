package entity

type Ticket struct {
	ID           int     `json:"ticket_id,omitempty" db:"id_ticket"`
	UserID       int     `json:"user_id,omitempty" db:"id_user"`
	ImagePrice   string  `json:"image_price,omitempty" db:"image_price"`
	NameProduct  string  `json:"name_product,omitempty" db:"name_product"`
	PriceProduct float64 `json:"price_product,omitempty" db:"price_product"`
	Status       int     `json:"status,omitempty" db:"status"`
}

const (
	TicketStatusOpen   int = 1
	TicketStatusClose  int = 2
	TicketStatusCancel int = 3
)

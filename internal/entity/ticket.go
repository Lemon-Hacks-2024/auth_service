package entity

import "time"

type TicketInput struct {
	//PhotoURL     string `json:"photo_url"`
	PriceProduct float64 `json:"price_product"`
	NameProduct  string  `json:"name_product"`
	//StoreAdress  string  `json:"store_adress"`
}

type Ticket struct {
	ID                 int    `json:"id,omitempty" db:"id"`
	Status             int    `json:"status,omitempty" db:"status"`
	RefusedDescription string `json:"refused_description,omitempty" db:"refusedDescription"`
	DateUpdate         string `json:"date_update,omitempty" db:"dateUpdate"`
	DateCreate         string `json:"date_create,omitempty" db:"dateCreate"`
}

type TicketDetails struct {
	ID              int    `json:"id,omitempty" db:"id"`
	TicketID        int    `json:"ticket_id,omitempty" db:"ticketId"`
	ReportedUser    int    `json:"reported_user,omitempty" db:"reportedUser"`
	TicketCity      string `json:"ticket_city,omitempty" db:"ticketCity"`
	SellerName      string `json:"seller_name,omitempty" db:"sellerName"`
	StoreAddress    string `json:"store_address,omitempty" db:"storeAddress"`
	ItemName        string `json:"item_name,omitempty" db:"itemName"`
	RecommededPrice string `json:"recommeded_price,omitempty" db:"recommededPrice"`
	CurrentPrice    string `json:"current_price,omitempty" db:"currentPrice"`
}

type TicketPhotos struct {
	ID         int       `json:"id,omitempty" db:"id"`
	TicketID   int       `json:"ticket_id,omitempty" db:"ticketId"`
	PhotoURL   string    `json:"photo_url,omitempty" db:"photoUrl"`
	Type       string    `json:"type,omitempty" db:"type"`
	DateUpdate time.Time `json:"date_update,omitempty" db:"dateUpdate"`
	DateCreate time.Time `json:"date_create,omitempty" db:"dateCreate"`
}

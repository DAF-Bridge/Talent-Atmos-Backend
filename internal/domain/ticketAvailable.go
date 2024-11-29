package domain

import "gorm.io/gorm"

type TicketAvailable struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	EventID     uint    `json:"event_id"`
	Event       Event   `json:"event"`
}

package model

import "time"

type Product struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	Available   int       `json:"available"`
	CreatedAt   time.Time `json:"created_at"`
}

package model

import "time"

type Product struct {
	SKU         string    `json:"sku"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	Available   int       `json:"available"`
	CreatedAt   time.Time `json:"created_at"`
	Eans        []string  `json:"eans"`
}

package model

type Product struct {
	SKU         string  `json:"sku"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Available   int     `json:"available"`
}

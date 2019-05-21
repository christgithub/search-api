package service

import . "github.com/search-api/model"

type Elasticer interface {
	Search(sku int) *Product
	Delete(sku int) bool
}

type Elastic struct {
	Port string `json:"port"`
}

func NewElastic() *Elastic {
	return &Elastic{
		Port: "9200",
	}
}

func Search(sku int) *Product {
	return nil
}

func Delete(sku int) bool {
	return true
}

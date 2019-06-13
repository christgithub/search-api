package service

import (
	"encoding/json"

	. "github.com/search-api/model"
	"github.com/search-api/repository"
)

type Elasticer interface {
	Search(sku string) (*Product, error)
	Delete(sku int) (bool, error)
	Add() (bool, error)
}

type Elastic struct {
	ElasticRepo repository.ElasticSearcher
}

func (e Elastic) Add() (bool, error) {
	return false, nil
}

func (e Elastic) Search(sku string) (*Product, error) {
	result, _ := e.ElasticRepo.SearchByID(sku)
	product := &Product{}
	json.Unmarshal(*result.Hits.Hits[0].Source, &product)

	return product, nil
}

func (e Elastic) Delete(sku int) (bool, error) {
	return false, nil
}

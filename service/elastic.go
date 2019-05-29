package service

import (
	. "github.com/search-api/model"
	"github.com/search-api/repository"
)

type Elasticer interface {
	Search(sku int) (*Product, error)
	Delete(sku int) (bool, error)
	Add() (bool, error)
}

type Elastic struct {
	ElasticRepo repository.ElasticSearcher
}

func (e Elastic) Add() (bool, error) {
	return false, nil
}

func (e Elastic) Search(sku int) (*Product, error) {
	_, _ = e.ElasticRepo.SearchByID(sku)
	return nil, nil
}

func (e Elastic) Delete(sku int) (bool, error) {
	return false, nil
}

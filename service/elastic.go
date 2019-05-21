package service

import (
	"github.com/olivere/elastic"
	. "github.com/search-api/model"
)

type Elasticer interface {
	Search(sku int) (*Product, error)
	Delete(sku int) (bool, error)
	Add() (bool, error)
}

type Elastic struct {
	ElasticClient *elastic.Client
}

func NewElastic() (*Elastic, error) {
	clientElastic, err := elastic.NewClient()

	if err != nil {
		return nil, err
	}

	return &Elastic{
		ElasticClient: clientElastic,
	}, nil
}

func (e Elastic) Add() (bool, error) {
	return false, nil
}

func (e Elastic) Search(sku int) (*Product, error) {
	return nil, nil
}

func (e Elastic) Delete(sku int) (bool, error) {
	return false, nil
}

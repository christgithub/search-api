package service

import (
	"time"

	. "github.com/search-api/model"
)

type ElasticStub struct{}

func (e ElasticStub) Add() (bool, error) {
	return false, nil
}

func (e ElasticStub) Search(sku int) (*Product, error) {
	p := &Product{
		"1",
		"Beans",
		2.99,
		1,
		time.Date(2019, 05, 18, 12, 34, 15, 651387237, time.UTC),
	}

	return p, nil
}

func (e ElasticStub) Delete(sku int) (bool, error) {
	return false, nil
}

package service

import (
	"time"

	. "github.com/search-api/model"
)

type ElasticMock struct{}

func (e ElasticMock) Add() (bool, error) {
	return false, nil
}

func (e ElasticMock) Search(sku int) (*Product, error) {
	createAt := time.Date(2019, 05, 18, 12, 34, 15, 651387237, time.UTC)

	return &Product{
		"1",
		"Beans",
		2.99,
		1,
		createAt,
	}, nil
}

func (e ElasticMock) Delete(sku int) (bool, error) {
	return false, nil
}

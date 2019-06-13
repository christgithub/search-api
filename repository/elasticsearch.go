package repository

import (
	"context"

	. "gopkg.in/olivere/elastic.v5"
)

type ElasticSearcher interface {
	SearchByID(sku string) (*SearchResult, error)
}

type ElasticSearch struct {
	Client Client
}

func (e ElasticSearch) SearchByID(sku string) (*SearchResult, error) {

	ctx := context.Background()
	termQuery := NewTermQuery("id", sku)
	result, err := e.Client.Search().
		Index("products").
		Query(termQuery).
		From(0).Size(1).
		Do(ctx)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (e ElasticSearch) SearchByEan(ean string) (*SearchResult, error) {

	ctx := context.Background()
	termQuery := NewTermQuery("ean", ean)
	result, err := e.Client.Search().
		Index("products").
		Query(termQuery).
		From(0).Size(1).
		Do(ctx)

	if err != nil {
		return nil, err
	}
	return result, nil
}

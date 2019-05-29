package repository

import (
	"context"

	"github.com/lytics/logrus"
	. "github.com/olivere/elastic"
)

type ElasticSearcher interface {
	SearchByID(sku int) (*GetResult, error)
}

type ElasticSearch struct {
	Client Client
}

func (e ElasticSearch) SearchByID(sku int) (*GetResult, error) {

	ctx := context.Background()

	product, err := e.Client.Get().
		Index("products").
		Type("product").
		Id("1").Do(ctx)

	if err != nil {
		return nil, err
	}

	if product.Found {
		logrus.Printf("Got document %s in version %d from index %s, type %s\n",
			product.Id,
			product.Version,
			product.Index,
			product.Type)
	}

	return product, nil
}

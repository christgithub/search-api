package service

import (
	"testing"

	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/prometheus/common/log"
	"github.com/search-api/repository/repositoryfakes"
	. "gopkg.in/olivere/elastic.v5"
)

type FakeSource []byte

func TestSearchApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Elastic Search Suite")
}

var _ = Describe("Elastic Search Suite", func() {
	Context("When a product is searched by sku", func() {
		It("returns a product", func() {

			fakeClient := &repositoryfakes.FakeElasticSearcher{}
			sku := "1"
			elasticService := Elastic{
				ElasticRepo: fakeClient,
			}

			rawMessage := &json.RawMessage{}
			source := []byte(`{"sku":"1"}`)
			err := json.Unmarshal(source, rawMessage)

			if err != nil {
				log.Fatal("Couldn't unmarshal product source")
			}

			searchHits := &SearchHits{
				Hits: []*SearchHit{
					{Source: rawMessage},
				},
			}

			fakeResult := &SearchResult{
				Hits: searchHits,
			}

			fakeClient.SearchByIDReturns(fakeResult, nil)
			productFromResult, _ := elasticService.Search(sku)

			Expect(fakeClient.SearchByIDCallCount()).To(Equal(1))
			Expect(productFromResult.SKU).To(Equal(sku))
		})
	})

	Context("When a product is searched by ean", func() {
		It("returns a product", func() {

			fakeClient := &repositoryfakes.FakeElasticSearcher{}
			sku := "1"
			elasticService := Elastic{
				ElasticRepo: fakeClient,
			}

			rawMessage := &json.RawMessage{}
			source := []byte(`{"sku":"1"}`)
			err := json.Unmarshal(source, rawMessage)

			if err != nil {
				log.Fatal("Couldn't unmarshal product source")
			}

			searchHits := &SearchHits{
				Hits: []*SearchHit{
					{Source: rawMessage},
				},
			}

			fakeResult := &SearchResult{
				Hits: searchHits,
			}

			fakeClient.SearchByIDReturns(fakeResult, nil)
			productFromResult, _ := elasticService.Search(sku)

			Expect(fakeClient.SearchByIDCallCount()).To(Equal(1))
			Expect(productFromResult.SKU).To(Equal(sku))
		})
	})
})

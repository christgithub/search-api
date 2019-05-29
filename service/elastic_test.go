package service

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/search-api/repository/repositoryfakes"
)

func TestSearchApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Elastic Search Suite")
}

var _ = Describe("Elastic Search Suite", func() {
	Context("When a product is search by id", func() {
		It("returns an elastic result", func() {

			fakeClient := &repositoryfakes.FakeElasticSearcher{}
			sku := 1
			elasticService := Elastic{
				ElasticRepo: fakeClient,
			}

			_, _ = elasticService.Search(sku)

			Expect(fakeClient.SearchByIDCallCount()).To(Equal(1))
			//Expect(product.ID).To(Equal(1))
		})
	})
})

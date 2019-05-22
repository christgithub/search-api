package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/search-api/service"

	"github.com/julienschmidt/httprouter"

	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/search-api/model"
	"github.com/search-api/service/servicefakes"
)

func TestSearchApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api Search Suite")
}

var _ = Describe("Search Api", func() {
	Context("when health endpoint is called", func() {
		It("returns an ok message", func() {
			h := &service.Handlers{
				Elastic: service.ElasticStub{},
			}
			server, _ := NewServer(h)
			req, _ := http.NewRequest("GET", "/health", nil)
			resp := httptest.NewRecorder()
			server.health(resp, req, httprouter.Params{})
			Expect(resp.Body.String()).To(Equal("test ok\n"))
		})

		It("returns an http ok 200", func() {

			h := &service.Handlers{
				Elastic: &servicefakes.FakeElasticer{},
			}

			server, _ := NewServer(h)
			req, _ := http.NewRequest("GET", "/health", nil)
			resp := httptest.NewRecorder()
			server.health(resp, req, httprouter.Params{})
			Expect(resp.Code).To(Equal(200))
		})
	})

	Context("when the search endpoint is called", func() {
		It("returns a product", func() {
			prd := &model.Product{
				"1",
				"Beans",
				2.99,
				1,
				time.Date(2019, 05, 18, 12, 34, 15, 651387237, time.UTC),
			}

			var fakeElastic = &servicefakes.FakeElasticer{}
			fakeElastic.SearchReturns(prd, nil)
			h := &service.Handlers{
				Elastic: fakeElastic,
			}

			server, _ := NewServer(h)
			r, _ := http.NewRequest("GET", "/products/1", nil)
			w := httptest.NewRecorder()
			p := httprouter.Params{}
			server.search(w, r, p)

			Expect(w.Body.String()).To(Equal(`{"id":"1","description":"Beans","price":2.99,"available":1,"created_at":"2019-05-18T12:34:15.651387237Z"}`))
			Expect(fakeElastic.SearchCallCount()).To(Equal(1))
		})
	})
})

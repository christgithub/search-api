package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"time"

	"fmt"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/search-api/handlers"
	"github.com/search-api/model"
	"github.com/search-api/service/servicefakes"
)

func TestSearchApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api Search Suite")
}

var _ = Describe("Search Api", func() {

	var server Server
	var fakeElastic = &servicefakes.FakeElasticer{}
	var h = &handlers.SearchHandler{
		Elastic: fakeElastic,
	}

	BeforeEach(func() {
		server, _ = NewServer(h)
		server.Routes()
	})

	Context("when health endpoint is called", func() {
		It("returns an ok message", func() {
			req, _ := http.NewRequest("GET", "/health", nil)
			resp := httptest.NewRecorder()
			server.Health(resp, req)
			Expect(resp.Body.String()).To(Equal("ok"))
		})

		It("returns an http ok 200", func() {
			req, _ := http.NewRequest("GET", "/health", nil)
			resp := httptest.NewRecorder()
			server.Health(resp, req)
			Expect(resp.Code).To(Equal(200))
		})
	})

	Context("when the search endpoint is called", func() {
		It("returns a product", func() {
			prd := &model.Product{
				ID:          "1",
				Description: "Beans",
				Price:       2.99,
				Available:   1,
				CreatedAt:   time.Date(2019, 05, 18, 12, 34, 15, 651387237, time.UTC),
			}

			fakeElastic.SearchReturns(prd, nil)
			w := httptest.NewRecorder()

			id := "1"
			endpoint := fmt.Sprintf("/search/id/%s", id)
			method := "GET"

			req, _ := http.NewRequest(method, endpoint, nil)
			req = mux.SetURLVars(req, map[string]string{"id": "1"})

			server.SearchHandler.ServeHTTP(w, req)

			Expect(w.Body.String()).To(Equal(`{"id":"1","description":"Beans","price":2.99,"available":1,"created_at":"2019-05-18T12:34:15.651387237Z"}`))
			Expect(fakeElastic.SearchCallCount()).To(Equal(1))
		})
	})
})

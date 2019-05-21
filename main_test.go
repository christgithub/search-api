package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/search-api/service"

	"github.com/julienschmidt/httprouter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSearchApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api Search Suite")
}

var _ = Describe("Search Api", func() {
	Context("when health endpoint is called", func() {
		It("returns an ok message", func() {
			h := &service.Handlers{
				service.ElasticMock{},
			}
			server, _ := NewServer(h)
			req, _ := http.NewRequest("GET", "/health", nil)
			resp := httptest.NewRecorder()
			server.health(resp, req, httprouter.Params{})
			Expect(resp.Body.String()).To(Equal("test ok\n"))
		})

		It("returns an http ok 200", func() {
			server, _ := NewServer()
			req, _ := http.NewRequest("GET", "/health", nil)
			resp := httptest.NewRecorder()
			server.health(resp, req, httprouter.Params{})
			Expect(resp.Code).To(Equal(200))
		})
	})

	Context("when the search endpoint is called", func() {
		It("returns a product", func() {
			server, _ := NewServer()
			r, _ := http.NewRequest("GET", "/products", nil)
			w := httptest.NewRecorder()
			server.search(w, r, httprouter.Params{})
			Expect(w.Body.String()).To(Equal(`{"result": "product"}`))
		})
	})
})

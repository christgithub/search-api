package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
			server := NewServer()
			req, _ := http.NewRequest("GET", "/health", nil)
			resp := httptest.NewRecorder()
			server.health(resp, req, httprouter.Params{})
			Expect(resp.Body.String()).To(Equal("test ok\n"))
		})

		It("returns an http ok 200", func() {
			server := NewServer()
			req, _ := http.NewRequest("GET", "/health", nil)
			resp := httptest.NewRecorder()
			server.health(resp, req, httprouter.Params{})
			Expect(resp.Code).To(Equal(200))
		})
	})
})

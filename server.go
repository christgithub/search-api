package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/search-api/handlers"
)

type Server struct {
	Router        *mux.Router
	SearchHandler *handlers.SearchHandlerBySku
	Port          string
}

func NewServer(h *handlers.SearchHandlerBySku) (Server, error) {
	server := Server{
		Router:        mux.NewRouter(),
		SearchHandler: h,
		Port:          "8000",
	}
	return server, nil
}

func (s *Server) Routes() {
	s.Router.Handle("/search/sku/{sku}", s.SearchHandler).Methods("GET")
	s.Router.Handle("/search/ean/{ean}", s.SearchHandler).Methods("GET")
	s.Router.HandleFunc("/health", s.Health).Methods("GET")

}

func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func (s *Server) Run() {
	address := fmt.Sprintf(":%s", s.Port)
	http.ListenAndServe(address, s.Router)
}

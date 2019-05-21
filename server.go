package main

import (
	"fmt"
	"net/http"

	"github.com/search-api/service"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	Router   *httprouter.Router
	Handlers *service.Handlers
	Port     string
}

func NewServer(h *service.Handlers) (*Server, error) {
	server := &Server{
		Router:   httprouter.New(),
		Handlers: h,
		Port:     "8000",
	}
	server.routes()
	return server, nil
}

func (s Server) routes() {
	s.Router.GET("/health", s.health)
	s.Router.GET("/search/:id", s.search)
}

func (s Server) Run() {
	address := fmt.Sprintf(":%s", s.Port)
	http.ListenAndServe(address, s.Router)
}

func (s Server) health(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	s.Handlers.Health(w, r, params)
}

func (s Server) search(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	s.Handlers.Search(w, r, params)
}

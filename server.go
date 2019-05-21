package main

import (
	"fmt"
	"net/http"

	"github.com/search-api/service"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	Router         *httprouter.Router
	ElasticService *service.Elastic
	Port           string
}

func NewServer() *Server {
	server := &Server{
		ElasticService: service.NewElastic(),
		Router:         httprouter.New(),
		Port:           "8000",
	}
	server.routes()
	return server
}

func (s Server) routes() {
	s.Router.GET("/health", s.health)
}

func (s Server) Run() {
	address := fmt.Sprintf(":%s", s.Port)
	http.ListenAndServe(address, s.Router)
}

func (s Server) health(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	service.Health(w, r, params)
}

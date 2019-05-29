package main

import (
	handlers2 "github.com/search-api/handlers"
	"github.com/search-api/repository"
	"github.com/search-api/service"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logger = logrus.New()
}

func main() {
	elastic := service.Elastic{
		ElasticRepo: &repository.ElasticSearch{},
	}

	handlers := &handlers2.SearchHandler{
		Elastic: elastic,
	}

	server, err := NewServer(handlers)
	if err != nil {
		logger.Fatal(err)
	}
	server.Routes()
	server.Run()
}

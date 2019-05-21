package main

import (
	"github.com/search-api/service"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	logger = logrus.New()
}

func main() {
	elastic, err := service.NewElastic()

	if err != nil {
		logger.Fatal(err)
	}

	handlers := &service.Handlers{
		*elastic,
	}

	server, err := NewServer(handlers)
	if err != nil {
		logger.Fatal(err)
	}
	server.Run()
}

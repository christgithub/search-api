package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/search-api/service"
	"github.com/sirupsen/logrus"
)

type SearchHandlerBySku struct {
	Elastic Elasticer
}

func (h SearchHandlerBySku) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	id := vars["sku"]

	product, err := h.Elastic.Search(id)

	if err != nil {

		logrus.Fatal(err)
	}

	productJson, err := json.Marshal(product)

	if err != nil {
		logrus.Fatal(err)
	}
	w.Write(productJson)
}

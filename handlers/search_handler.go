package handlers

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
	. "github.com/search-api/service"
	"github.com/sirupsen/logrus"
)

type SearchHandler struct {
	Elastic Elasticer
}

func (h SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		logrus.Fatal(err)
	}

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

package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/julienschmidt/httprouter"
)

type Handlers struct {
	Elastic Elastic
}

func (h Handlers) Health(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "test ok\n")
}

func (h Handlers) Search(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	product_id := params.ByName("id")
	id, err := strconv.Atoi(product_id)

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

	fmt.Fprint(w, string(productJson))
}

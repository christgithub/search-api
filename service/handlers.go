package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/julienschmidt/httprouter"
)

type Handlers struct {
	Elastic Elasticer
}

func (h Handlers) Health(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "test ok\n")
}

func (h Handlers) Search(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	//product_id := params.ByName("id")
	//fmt.Printf("PRODUCT ID %v", product_id)
	//id, err := strconv.Atoi(product_id)
	//
	//if err != nil {
	//	logrus.Fatal(err)
	//}

	product, err := h.Elastic.Search(1)

	if err != nil {
		logrus.Fatal(err)
	}

	productJson, err := json.Marshal(product)

	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Fprint(w, string(productJson))
}

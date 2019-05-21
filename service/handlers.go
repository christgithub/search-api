package service

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Health(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "test ok\n")
}

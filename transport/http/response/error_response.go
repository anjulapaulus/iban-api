package response

import (
	"log"
	"net/http"

	"github.com/anjulapaulus/iban-api/transport/http/middleware/errors"
	requestErrors "github.com/anjulapaulus/iban-api/transport/http/request/errors"
)

func Error(w http.ResponseWriter, err interface{}) {
	log.Default().Print(err)

	if _, ok := err.(*errors.MiddlewareError); ok {
		Send(w, http.StatusBadRequest, err)
	}

	if e, ok := err.(*requestErrors.ValidationError); ok {
		Send(w, http.StatusBadRequest, e)
	}
}

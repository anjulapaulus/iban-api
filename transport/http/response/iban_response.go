package response

import (
	"net/http"
)

type IBANResponse struct {
	Valid bool `json:"valid"`
}

func EncodeResponse(w http.ResponseWriter, r interface{}) error {

	if r == nil {
		req := IBANResponse{}
		Send(w, http.StatusOK, req)
		return nil
	}
	Send(w, http.StatusOK, r)

	return nil
}

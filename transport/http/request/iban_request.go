package request

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/anjulapaulus/iban-api/transport/http/request/errors"
)

type IbanRequest struct {
	IBAN string `json:"iban"`
}

func DecodeRequest(ctx context.Context, req *http.Request) (*IbanRequest, error) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, errors.NewValidationError(
			"1001",
			"Error reading request body.",
			err,
		)
	}

	request := IbanRequest{}

	err = json.Unmarshal(data, &request)
	if err != nil {
		return nil, errors.NewValidationError(
			"1002",
			"Error unmarshalling request body - No request body",
			err,
		)
	}

	err = request.validate()
	if err != nil {
		return nil, err
	}

	return &request, nil

}

func (r *IbanRequest) validate() error {
	if r.IBAN == "" {
		return errors.NewValidationError(
			"1003",
			"Request validation error | IBAN passed is empty",
			nil,
		)
	}
	return nil
}

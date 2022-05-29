package controllers

import (
	"net/http"

	"github.com/anjulapaulus/iban-api/domain/usecases"
	"github.com/anjulapaulus/iban-api/iban"
	"github.com/anjulapaulus/iban-api/transport/http/request"
	"github.com/anjulapaulus/iban-api/transport/http/response"
)

// IBANController contains controller logic for IBAN endpoint
type IBANController struct {
	usecase *usecases.IBANUsecase
}

func NewIBANController() *IBANController {

	validator, err := iban.NewIbanValidator(iban.CountryIBAN)
	if err != nil {
		panic(err)
	}
	// init and bind use cases
	usecase := &usecases.IBANUsecase{
		Validator: validator,
	}

	return &IBANController{
		usecase: usecase,
	}
}

// CheckIBAN validates a IBAN string
func (c *IBANController) CheckIBAN(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	request, err := request.DecodeRequest(ctx, r)
	if err != nil {
		response.Error(w, err)
		return
	}

	res, err := c.usecase.ValidateIBAN(ctx, request)
	if err != nil {
		response.Error(w, err)
		return
	}

	err = response.EncodeResponse(w, res)
	if err != nil {
		response.Error(w, err)
		return
	}
}

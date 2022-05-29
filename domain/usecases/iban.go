package usecases

import (
	"context"

	"github.com/anjulapaulus/iban-api/iban"
	"github.com/anjulapaulus/iban-api/transport/http/request"
	"github.com/anjulapaulus/iban-api/transport/http/response"
)

type IBANUsecase struct {
	Validator *iban.IbanValidator
}

func (u *IBANUsecase) ValidateIBAN(ctx context.Context, req *request.IbanRequest) (interface{}, error) {
	check, _ := u.Validator.Validate(req.IBAN)
	// if err != nil {
	// 	return nil, err
	// }

	return response.IBANResponse{
		Valid: check,
	}, nil
}

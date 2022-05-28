package iban

import (
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type IbanValidator struct {
	countryIBANS map[string]Country
	rwMu         *sync.RWMutex
}

func NewIbanValidator(cinfo map[string]Country) (*IbanValidator, error) {
	if len(cinfo) < 1 {
		return nil, errors.New("nil map passed")
	}
	return &IbanValidator{
		countryIBANS: cinfo,
		rwMu:         new(sync.RWMutex),
	}, nil
}

func (v *IbanValidator) Validate(iban string) (bool, error) {
	// convert string to uppercase
	s := strings.ToUpper(iban)

	// remove spaces in string
	s = strings.Replace(s, " ", "", -1)

	// extract country code
	countryCode := s[0:2]

	// Get country settings for country code
	v.rwMu.RLock()
	defer v.rwMu.RUnlock()
	ci, ok := v.countryIBANS[countryCode]
	if !ok {
		return false, fmt.Errorf("unsupported country code %v", countryCode)
	}

	// validate total iban length based on country
	if len(s) != ci.Length {
		return false, fmt.Errorf("IBAN length and length specified cofr country code %v mismatch. Actual IBAN length: %v, Specified length: %v", countryCode, len(s), ci.Length)
	}

	// Move the four initial characters to the end of the string
	rearrangedIban := s[4:] + s[:4]

	// check if characters are A-z
	regx := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

	intStr := ""
	for _, c := range rearrangedIban {
		if regx(string(c)) {
			asciiValue := int(c)
			asciiValue -= 55
			intStr += strconv.Itoa(asciiValue)
			continue
		}
		intStr += string(c)
	}

	n := new(big.Int)
	n, ok = n.SetString(intStr, 10)
	if !ok {
		return false, errors.New("SetString error")
	}
	divisor := new(big.Int).SetInt64(97)

	remainder := new(big.Int).Mod(n, divisor)

	if remainder.Int64() == 1 {
		return true, nil
	}

	return false, errors.New("IBAN has incorrect check digits")
}

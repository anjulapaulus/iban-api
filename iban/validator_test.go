package iban

import "testing"

func TestNewIbanValidator(t *testing.T) {
	validator, err := NewIbanValidator(CountryIBAN)
	if err != nil {
		t.Error(err)
	}

	if validator == nil {
		t.Error("validator is nil")
	}
}

func TestValidateFakeIbans(t *testing.T) {
	validator, err := NewIbanValidator(CountryIBAN)
	if err != nil {
		t.Error(err)
	}

	ibans := []string{
		"VG96VPVG00000L2345678901",
		"1234567890",
		"12345678901234567890",
		"NL30ABNA0123456789",
		"NL30ABNA0517552264",
		"NL30ABNA05175522AB",
	}
	for _, v := range ibans {
		_, err := validator.Validate(v)
		if err == nil {
			t.Error("test error: cannot identify wrong ibans")
		}
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		ibans  []string
		result bool
	}{
		{
			ibans: []string{
				"AE020200000030124176201",
				"AE07 0331 2345 6789 0123 456",
				"AE14 0340 0000 1401 1019 050",
				"AE260211000000230064016",
				"AE320030000100228001001",
				"AE320030010274073001001",
				"AE320330000010195510887",
				"AE730030000789456123456",
				"AE940350000000250008661",
				"AE950260001261025056501",
			},
			result: true,
		},
		{
			ibans: []string{
				"AL 79208110080000001043631801",
				"AL 45209516090000600388970102",
				"AL85 2021 1037 0000 0000 0620 5792",
				"AL05 2141 1144 0111 2930 4302 0418",
				"AL38202110130000000000157945",
				"AL35202110130000000002157945",
				"AL11 2051 1014 3288 86CL TJCF EURA",
				"AL 4920 5110 1427 3430 CLTJCLALLA",
				"AL81202110130000000021002859",
				"AL 05 2121 1016 0000 0000 0028 8405",
			},
			result: true,
		},
	}

	validator, err := NewIbanValidator(CountryIBAN)
	if err != nil {
		t.Error(err)
	}

	for _, tt := range tests {
		for _, v := range tt.ibans {
			check, err := validator.Validate(v)
			if check != tt.result {
				t.Error("Failed: validate test ", v)
			}
			if err != nil {
				t.Error(err)
			}

		}
	}

}

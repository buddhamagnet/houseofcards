package validators_test

import (
	"testing"

	. "github.com/buddhamagnet/houseofcards/validators"
)

func TestCheckPresenceSuccess(t *testing.T) {
	code, _ := CheckPresent.Validate("1234")
	if code != 200 {
		t.Fatalf("expected presence check to return 200, got %d", code)
	}
}

func TestCheckPresenceFail(t *testing.T) {
	code, _ := CheckPresent.Validate("")
	if code != 400 {
		t.Fatalf("expected presence check to return 400, got %d", code)
	}
}

func TestCheckDigitsSuccess(t *testing.T) {
	code, _ := CheckDigits.Validate("1234")
	if code != 200 {
		t.Fatalf("expected digit check to return 200, got %d", code)
	}
}

func TestCheckDigitsFail(t *testing.T) {
	code, _ := CheckDigits.Validate("DAVE")
	if code != 400 {
		t.Fatalf("expected digit check to return 400, got %d", code)
	}
}

func TestCheckDigitLengthSuccess(t *testing.T) {
	validator := NewDigitLengthValidator(4)
	code, _ := validator.Validate("1234")
	if code != 200 {
		t.Fatalf("expected length check to return 200, got %d", code)
	}
}

func TestCheckDigitLengthFail(t *testing.T) {
	validator := NewDigitLengthValidator(4)
	code, _ := validator.Validate("12345")
	if code != 400 {
		t.Fatalf("expected length check to return 400, got %d", code)
	}
	code, _ = validator.Validate("DAVE")
	if code != 400 {
		t.Fatalf("expected presence check to return 400, got %d", code)
	}
}

func TestCheckAccordHotelsSuccess(t *testing.T) {
	validator := NewAccordHotelsValidator()
	code, _ := validator.Validate("1234567890")
	if code != 200 {
		t.Fatalf("expected Accord Hotels check to return 200, got %d", code)
	}
}

func TestCheckAccordHotelsFail(t *testing.T) {
	validator := NewAccordHotelsValidator()
	code, _ := validator.Validate("1234")
	if code != 400 {
		t.Fatalf("expected Accord Hotels check to return 200, got %d", code)
	}
}

func TestCheckMilesAndMoreSuccess(t *testing.T) {
	validator := NewMilesAndMoreValidator()
	code, _ := validator.Validate("123456789012345")
	if code != 200 {
		t.Fatalf("expected Miles And More check to return 200, got %d", code)
	}
}

func TestCheckMilesAndMoreFail(t *testing.T) {
	validator := NewMilesAndMoreValidator()
	code, _ := validator.Validate("1234")
	if code != 400 {
		t.Fatalf("expected Miles And More check to return 200, got %d", code)
	}
}

func TestCheckTajHotelsSuccess(t *testing.T) {
	validator := NewTajHotelsValidator()
	code, _ := validator.Validate("10101222222222")
	if code != 200 {
		t.Fatalf("expected Taj Hotels check to return 200, got %d", code)
	}
}

func TestCheckTajHotelsFail(t *testing.T) {
	validator := NewTajHotelsValidator()
	code, _ := validator.Validate("10121222222222")
	if code != 400 {
		t.Fatalf("expected Taj Hotels check to return 200, got %d", code)
	}
}

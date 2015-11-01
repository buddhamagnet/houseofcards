package validators

import (
	"fmt"
	"net/http"
	"regexp"
)

var (
	// CheckPresent validates presence of the code.
	CheckPresent PartnerCodeValidator
	// CheckDigits validates numericality of the code.
	CheckDigits PartnerCodeValidator
	// Map maps partner names to validators.
	Map map[string]PartnerCodeValidator
)

func init() {
	CheckPresent = new(PresenceValidator)
	CheckDigits = NewDigitValidator()

	Map = map[string]PartnerCodeValidator{
		"cathaypacific":     CheckPresent,
		"singapore":         CheckPresent,
		"jetairways":        CheckPresent,
		"airasia":           CheckPresent,
		"jetairwaysindia":   CheckDigits,
		"finnair":           CheckDigits,
		"pins":              CheckDigits,
		"tapvictoria":       CheckDigits,
		"saseurobonus":      CheckDigits,
		"iberia":            CheckDigits,
		"etihad":            CheckDigits,
		"ethiopianairlines": CheckDigits,
		"accordhotels":      NewAccordHotelsValidator(),
		"milesandmore":      NewMilesAndMoreValidator(),
		"tajhotels":         NewTajHotelsValidator(),
	}
}

// PartnerCodeValidator is the main validation interface.
type PartnerCodeValidator interface {
	Validate(card string) (int, string)
}

// PresenceValidator validates the presence of the number.
type PresenceValidator struct{}

// Validate does the actual validation.
func (p *PresenceValidator) Validate(card string) (int, string) {
	if card != "" {
		return http.StatusOK, ""
	}
	return http.StatusBadRequest, "missing code"
}

// NewDigitValidator returns a digit validator.
func NewDigitValidator() *PatternValidator {
	digits := regexp.MustCompile("^[0-9]+$")
	return &PatternValidator{pattern: digits, message: "invalid check digit"}
}

// NewDigitLengthValidator returns a digit length validator.
func NewDigitLengthValidator(l int) *PatternValidator {
	digitLength := regexp.MustCompile(fmt.Sprintf("^[0-9]{%d}$", l))
	return &PatternValidator{pattern: digitLength, message: "invalid length"}
}

// PatternValidator validates a number pattern.
type PatternValidator struct {
	pattern *regexp.Regexp
	message string
}

// NewPatternValidator returns a pattern validator.
func NewPatternValidator(pat string) *PatternValidator {
	return &PatternValidator{pattern: regexp.MustCompile(pat), message: "invalid code"}
}

// Validate does the actual validation.
func (p *PatternValidator) Validate(card string) (int, string) {
	if p.pattern.MatchString(card) {
		return http.StatusOK, ""
	}
	return http.StatusBadRequest, p.message
}

// NewAccordHotelsValidator returns a custom validator for Accord Hotels.
func NewAccordHotelsValidator() *PatternValidator {
	return NewDigitLengthValidator(10)
}

// NewMilesAndMoreValidator returns a custom validator for Miles And More.
func NewMilesAndMoreValidator() *PatternValidator {
	return NewDigitLengthValidator(15)
}

// NewTajHotelsValidator returns a custom validator for New Taj Hotels.
func NewTajHotelsValidator() *PatternValidator {
	return NewPatternValidator("^10101([0-9]{9}$)")
}

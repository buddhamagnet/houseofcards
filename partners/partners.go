package partners

import (
	"fmt"
	"net/http"
	"regexp"
)

var (
	presenceValidator PartnerCodeValidator
	digitValidator    PartnerCodeValidator
	// Map maps partner names to validators.
	Map map[string]PartnerCodeValidator
)

func init() {
	presenceValidator = new(PresenceValidator)
	digitValidator = NewDigitValidator()

	Map = map[string]PartnerCodeValidator{
		"cathaypacific":     presenceValidator,
		"singapore":         presenceValidator,
		"jetairways":        presenceValidator,
		"airasia":           presenceValidator,
		"jetairwaysindia":   digitValidator,
		"finnair":           digitValidator,
		"pins":              digitValidator,
		"tapvictoria":       digitValidator,
		"saseurobonus":      digitValidator,
		"iberia":            digitValidator,
		"etihad":            digitValidator,
		"ethiopianairlines": digitValidator,
		"accordhotels":      NewDigitLengthValidator(10),
		"milesandmore":      NewDigitLengthValidator(15),
		"tajhotels":         NewPatternValidator("^10101([0-9]{9}$)"),
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
		return http.StatusOK, "OK"
	}
	return http.StatusBadRequest, "BAD"
}

// NewDigitValidator returns a digit validator.
func NewDigitValidator() *PatternValidator {
	digits := regexp.MustCompile("^[0-9]+$")
	return &PatternValidator{pattern: digits}
}

// NewDigitLengthValidator returns a digit length validator.
func NewDigitLengthValidator(l int) *PatternValidator {
	digitLength := regexp.MustCompile(fmt.Sprintf("^[0-9]{%d}$", l))
	return &PatternValidator{pattern: digitLength}
}

// PatternValidator validates a number pattern.
type PatternValidator struct {
	pattern *regexp.Regexp
}

// Validate does the actual validation.
func (p *PatternValidator) Validate(card string) (int, string) {
	if p.pattern.MatchString(card) {
		return http.StatusOK, "OK"
	}
	return http.StatusBadRequest, "BAD"
}

// NewPatternValidator returns a pattern validator.
func NewPatternValidator(pat string) *PatternValidator {
	return &PatternValidator{pattern: regexp.MustCompile(pat)}
}

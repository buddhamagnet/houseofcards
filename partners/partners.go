package partners

import (
	"fmt"
	"net/http"
	"regexp"
)

var (
	patternTajHotels  *regexp.Regexp
	presenceValidator PartnerCodeValidator
	digitValidator    PartnerCodeValidator
	Map               map[string]PartnerCodeValidator
)

func init() {
	patternTajHotels = regexp.MustCompile("^10101([0-9]{9}$)")
	presenceValidator = new(PresenceValidator)
	presenceValidator = NewDigitValidator()

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
		"tajhotels":         NewPatternValidator(patternTajHotels),
	}
}

type PartnerCodeValidator interface {
	Validate(card string) (int, string)
}

type PresenceValidator struct{}

func (p *PresenceValidator) Validate(card string) (int, string) {
	if card != "" {
		return http.StatusOK, "OK"
	}
	return http.StatusBadRequest, "BAD"
}

func NewDigitValidator() *PatternValidator {
	digits := regexp.MustCompile("^[0-9]+$")
	return &PatternValidator{pattern: digits}
}

func NewDigitLengthValidator(l int) *PatternValidator {
	digitLength := regexp.MustCompile(fmt.Sprintf("^[0-9]{%d}$", l))
	return &PatternValidator{pattern: digitLength}
}

type PatternValidator struct {
	pattern *regexp.Regexp
}

func (p *PatternValidator) Validate(card string) (int, string) {
	if p.pattern.MatchString(card) {
		return http.StatusOK, "OK"
	}
	return http.StatusBadRequest, "BAD"
}

func NewPatternValidator(pat *regexp.Regexp) *PatternValidator {
	return &PatternValidator{pattern: pat}
}

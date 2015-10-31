package partners

import "net/http"

var Map = map[string]PartnerCodeValidator{
	"singapore":     new(SingaporeAirlines),
	"cathaypacific": new(CathayPacificAirlines),
}

type PartnerCodeValidator interface {
	Validate(card string) (int, string)
}

type SingaporeAirlines struct{}

func (s *SingaporeAirlines) Validate(card string) (int, string) {
	return CardMandatory(card)
}

type CathayPacificAirlines struct{}

func (c *CathayPacificAirlines) Validate(card string) (int, string) {
	return CardMandatory(card)
}

func CardMandatory(card string) (int, string) {
	if card != "" {
		return http.StatusOK, "OK"
	}
	return http.StatusBadRequest, "BAD"
}

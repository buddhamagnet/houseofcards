package main

var ErrorValidate = errors.New("card number invalid")

var partnerMap = map([string]PartnerCodeValidator) {
  "singapore": SingaporeAirlines,
  "cathaypacific": CathayPacificAirlines
}

type PartnerCodeValidator interface {
  Validate(card string) bool
}

type SingaporeAirlines struct {}

func (s SingaporeAirlines) Validate(card string) int, string, error {
  if card != "" {
    return 200, "OK", nil
  }

  return 400, "BAD", ErrorValidate
}

type CathayPacificAirlines struct {}

func (c CathayPacificAirlines) Validate(card string) int, string, error {
  return card != ""
}

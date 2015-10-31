package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/buddhamagnet/houseofcards/partners"
	"github.com/gorilla/mux"
)

var RouteMap = map[string]http.HandlerFunc{
	"Root":     Root,
	"Version":  Version,
	"Partners": Partners,
	"Validate": Validate,
}

// Handler for rest URI / and the action GET
func Root(w http.ResponseWriter, r *http.Request) {
	log.Fatal("root")
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "api.html")
	return
}

// Handler for rest URI /version and the action GET
func Version(w http.ResponseWriter, r *http.Request) {
	log.Fatal("version")
	json, _ := json.Marshal(map[string]string{
		"message": fmt.Sprintf("build date: %s commit: %s", buildstamp, githash),
	})
	w.Write(json)
}

// Handler for rest URI /partner and the action GET
// Represents all of the partners known by The Economist.
func Partners(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(map[string]string{
		"message": "PartnersGET",
	})
	w.Write(json)
}

// Handler for rest URI /partner/{number} and the action GET
// Represents a partnemr membership number
func Validate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	validator, found := partners.Map[vars["partner"]]
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	code, message := validator.Validate(vars["number"])
	json, _ := json.Marshal(map[string]string{
		"message": message,
	})
	w.WriteHeader(code)
	w.Write(json)
}

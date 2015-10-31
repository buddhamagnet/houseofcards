package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/buddhamagnet/houseofcards/validators"
	"github.com/gorilla/mux"
)

// RouteMap links RAML display names to handlers.
var RouteMap = map[string]http.HandlerFunc{
	"Root":     Root,
	"Version":  Version,
	"Partners": Partners,
	"Validate": Validate,
}

// Root handler for rest URI / and the action GET
func Root(w http.ResponseWriter, r *http.Request) {
	log.Fatal("root")
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "api.html")
	return
}

// Version handler for rest URI /version and the action GET
func Version(w http.ResponseWriter, r *http.Request) {
	log.Fatal("version")
	json, _ := json.Marshal(map[string]string{
		"message": fmt.Sprintf("build date: %s commit: %s", buildstamp, githash),
	})
	w.Write(json)
}

// Partners handler for rest URI /partner and the action GET
func Partners(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(map[string]string{
		"message": "PartnersGET",
	})
	w.Write(json)
}

// Validate handler for rest URI /partners/{partner}/{number} and the action GET
func Validate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	validator, found := validators.Map[vars["partner"]]
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Validate the code.
	code, message := validator.Validate(vars["number"])
	response := map[string]string{
		"number": vars["number"],
	}
	// Populate the map based on validity.
	switch code {
	case http.StatusOK:
		response["location"] = fmt.Sprintf("/partners/%s/%s", vars["partner"], vars["number"])
	default:
		response["error"] = message
	}
	// Send it back down the wire.
	json, _ := json.Marshal(response)
	w.WriteHeader(code)
	w.Write(json)
}

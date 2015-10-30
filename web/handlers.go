package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var RouteMap = map[string]http.HandlerFunc{
	"Root":     Root,
	"Version":  Version,
	"Partners": Partners,
	"Validate": Validate,
}

// Handler for rest URI / and the action GET
func Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "api.html")
	return
}

// Handler for rest URI /version and the action GET
func Version(w http.ResponseWriter, r *http.Request) {
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
	json, _ := json.Marshal(map[string]string{
		"message": "ValidateGET",
	})
	w.Write(json)
}

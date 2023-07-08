package controllers

import (
	"encoding/json"
	"net/http"
)

type HealthController struct{}

func HealthStatus(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("OK")
}

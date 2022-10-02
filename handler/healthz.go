package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct {
	Endpoint string
}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{Endpoint: "/healthz"}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	healthResponse := &model.HealthzResponse{Message: "OK"}
	err := json.NewEncoder(w).Encode(healthResponse)
	if err != nil {
		log.Println(err)
	}
}

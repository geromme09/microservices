package handler

import (
	"encoding/json"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"
)

type HttpTripPreviewHandler struct {
	Service domain.TripService
}

func NewTripPreviewHandler(service domain.TripService) *HttpTripPreviewHandler {
	return &HttpTripPreviewHandler{
		Service: service,
	}
}

func (h *HttpTripPreviewHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {

	var req previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.Service.GetRoute(r.Context(), req.Pickup, req.Destination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := writeJsonResponse(w, http.StatusOK, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func writeJsonResponse(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return nil
}

type previewTripRequest struct {
	UserId      string           `json:"userId"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

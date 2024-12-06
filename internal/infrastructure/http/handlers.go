package http

import (
	"billing-sys/internal/application/usecases"
	"encoding/json"
	"net/http"
	"strconv"
)

type Handlers struct {
	// Use cases here
	CalculateChargeUC *usecases.CalculateChargeUseCase
}

func (h *Handlers) GetBuildingHandler(w http.ResponseWriter, r *http.Request) {
	reqID := r.Header.Get("id")
	_, err := strconv.Atoi(reqID)
	if err != nil {
		http.Error(w, "invalid request id: "+reqID, http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "this would call a usecase to get building details",
	})
}

func (h *Handlers) CreateBuildingHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name      string `json:"name"`
		Address   string `json:"address"`
		TotalArea string `json:"total_area"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "this would call a usecase to create a building details",
	})
}

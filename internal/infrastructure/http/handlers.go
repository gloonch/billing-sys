package http

import (
	"billing-sys/internal/application/dto"
	"billing-sys/internal/application/usecases/buildings"
	"billing-sys/internal/application/usecases/payments"
	"billing-sys/internal/application/usecases/units"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Handlers struct {
	// Use cases here
	CreateBuildingUC          *buildings.CreateBuildingUseCase
	GetBuildingUC             *buildings.GetBuildingUseCase
	ListBuildingsUC           *buildings.ListAllBuildingUseCase
	UpdateBuildingUC          *buildings.UpdateBuildingUseCase
	DeleteBuildingUC          *buildings.DeleteBuildingUseCase
	CalculateBuildingChargeUC *buildings.CalculateBuildingChargeUseCase

	// Unit use cases
	CreateUnitUC *units.CreateUnitUseCase
	GetUnitUC    *units.GetUnitUseCase
	ListUnitsUC  *units.ListAllUnitUseCase
	UpdateUnitUC *units.UpdateUnitUseCase
	DeleteUnitUC *units.DeleteUnitUseCase

	// Payment use cases
	CreatePaymentUC        *payments.CreatePaymentUseCase
	DeletePaymentUC        *payments.DeletePaymentUseCase
	ListPaymentsByUnitIDUC *payments.ListPaymentsByUnitIDUseCase
}

func (h *Handlers) CreateBuildingHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}

	var input dto.CreateBuildingInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)

		return
	}

	result, err := h.CreateBuildingUC.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handlers) GetBuildingHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}

	path := r.URL.Path
	segments := strings.Split(path, "/")

	if len(segments) < 3 || segments[1] != "buildings" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := segments[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)

		return
	}

	result, err := h.GetBuildingUC.Execute(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handlers) ListBuildingHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}

	result, err := h.ListBuildingsUC.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handlers) UpdateBuildingHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}

	path := r.URL.Path
	segments := strings.Split(path, "/")

	if len(segments) < 3 || segments[1] != "buildings" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := segments[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)

		return
	}

	var input dto.UpdateBuildingInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)

		return
	}

	result, err := h.UpdateBuildingUC.Execute(uint(id), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handlers) DeleteBuildingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	segments := strings.Split(path, "/")

	if len(segments) < 3 || segments[1] != "buildings" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := segments[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.DeleteBuildingUC.Execute(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Building deleted successfully",
	})
}

func (h *Handlers) CalculateBuildingChargeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")

	if len(segments) < 3 || segments[1] != "buildings" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := segments[2]
	buildingID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	charges, err := h.CalculateBuildingChargeUC.Execute(uint(buildingID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(charges)
}

// unit handlers TODO: should i separate these handlers into EntityHandler ?

func (h *Handlers) CreateUnitHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}

	var input dto.CreateUnitInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)

		return
	}

	result, err := h.CreateUnitUC.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

}

func (h *Handlers) GetUnitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")

	if len(segments) < 3 || segments[1] != "units" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := segments[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)

		return
	}

	result, err := h.GetUnitUC.Execute(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handlers) ListUnitHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}

	result, err := h.ListUnitsUC.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handlers) UpdateUnitHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")

	if len(segments) < 3 || segments[1] != "units" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := segments[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)

		return
	}

	var input dto.CreateUnitInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)

		return
	}

	result, err := h.UpdateUnitUC.Execute(uint(id), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handlers) DeleteUnitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")

	if len(segments) < 3 || segments[1] != "units" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := segments[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.DeleteUnitUC.Execute(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Unit deleted successfully",
	})
}

// payment handlers

func (h *Handlers) CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input dto.CreatePaymentInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	result, err := h.CreatePaymentUC.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handlers) ListPaymentsByUnitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")

	if len(segments) < 3 || segments[1] != "payments" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := segments[3]
	unitID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	results, err := h.ListPaymentsByUnitIDUC.Execute(uint(unitID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func (h *Handlers) DeletePaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")

	if len(segments) < 3 || segments[1] != "payments" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := segments[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.DeletePaymentUC.Execute(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Unit deleted successfully",
	})
}

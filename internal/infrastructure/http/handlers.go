package http

import (
	"billing-sys/internal/application/dto"
	"billing-sys/internal/application/usecases/buildings"
	"billing-sys/internal/application/usecases/payments"
	"billing-sys/internal/application/usecases/units"
	"billing-sys/internal/domain/strategies"
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

// CreateBuildingHandler godoc
// @Summary Create a new building
// @Description Adds a new building to the system
// @Tags buildings
// @Accept json
// @Produce json
// @Param building body dto.CreateBuildingInput true "Building data"
// @Success 201 {object} dto.CreateBuildingOutput
// @Failure 400 {string} string "Invalid input"
// @Failure 405 {string} string "Method Not Allowed"
// @Failure 500 {string} string "Internal Server Error"
// @Router /buildings [post]
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

// GetBuildingHandler godoc
// @Summary Get details of a specific building
// @Description Retrieves details of a building using its unique ID
// @Tags buildings
// @Accept json
// @Produce json
// @Param id path int true "Building ID"
// @Success 200 {object} dto.GetBuildingOutput
// @Failure 400 {string} string "Invalid URL or ID"
// @Failure 405 {string} string "Method Not Allowed"
// @Failure 500 {string} string "Internal Server Error"
// @Router /buildings/{id} [get]
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

// ListBuildingHandler godoc
// @Summary List all buildings
// @Description Retrieves a list of all buildings
// @Tags buildings
// @Accept json
// @Produce json
// @Success 200 {array} dto.GetBuildingOutput
// @Failure 405 {string} string "Method Not Allowed"
// @Failure 500 {string} string "Internal Server Error"
// @Router /buildings [get]
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

// UpdateBuildingHandler godoc
// @Summary Update a building's details
// @Description Updates the details of a specific building by its ID
// @Tags buildings
// @Accept json
// @Produce json
// @Param id path int true "Building ID"
// @Param building body dto.UpdateBuildingInput true "Updated building data"
// @Success 200 {object} dto.UpdateBuildingOutput
// @Failure 400 {string} string "Invalid URL, ID, or input data"
// @Failure 405 {string} string "Method Not Allowed"
// @Failure 500 {string} string "Internal Server Error"
// @Router /buildings/{id} [put]
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

// DeleteBuildingHandler godoc
// @Summary Delete a building
// @Description Deletes a building by its ID
// @Tags buildings
// @Accept json
// @Produce json
// @Param id path int true "Building ID"
// @Success 200 {object} map[string]string "Building deleted successfully"
// @Failure 400 {string} string "Invalid URL or ID"
// @Failure 405 {string} string "Method Not Allowed"
// @Failure 500 {string} string "Internal Server Error"
// @Router /buildings/{id} [delete]
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

// CalculateBuildingChargeHandler godoc
// @Summary Calculate charges for a building
// @Description Calculates the charges for a building using a specified strategy
// @Tags buildings
// @Accept json
// @Produce json
// @Param id path int true "Building ID"
// @Param strategy path int true "Charge calculation strategy: 1 (Area-Based), 2 (Occupant-Based), 3 (Combined)"
// @Success 200 {object} map[uint]float64 "Map of unit IDs to calculated charges"
// @Failure 400 {string} string "Invalid URL, ID, or strategy"
// @Failure 405 {string} string "Method Not Allowed"
// @Failure 500 {string} string "Internal Server Error"
// @Router /buildings/{id}/charges/{strategy} [get]
func (h *Handlers) CalculateBuildingChargeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")

	if len(segments) < 4 || segments[1] != "buildings" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := segments[2]
	buildingID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// pick strategy of charge calculation

	var strategy strategies.ChargeCalculationStrategy
	input_startegy, strgErr := strconv.Atoi(segments[4])
	if strgErr != nil {
		http.Error(w, "Invalid Strategy ", http.StatusBadRequest)

		return
	}
	if input_startegy > 3 || input_startegy < 0 {
		http.Error(w, "Invalid Strategy ", http.StatusBadRequest)

		return
	}
	if input_startegy == 1 {
		strategy = &strategies.AreaBasedStrategy{}
	}
	if input_startegy == 2 {
		strategy = &strategies.OccupantBasedStrategy{}
	}
	if input_startegy == 3 {
		strategy = &strategies.CombinedStrategy{}
	}

	charges, err := h.CalculateBuildingChargeUC.Execute(uint(buildingID), strategy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(charges)
}

// unit handlers TODO: should i separate these handlers into EntityHandler ?

// CreateUnitHandler godoc
// @Summary Create a new unit
// @Description Adds a new unit to the system
// @Tags units
// @Accept json
// @Produce json
// @Param unit body dto.CreateUnitInput true "Unit data"
// @Success 201 {object} dto.CreateUnitOutput
// @Failure 400 {string} string "Invalid input"
// @Failure 405 {string} string "Method Not Allowed"
// @Failure 500 {string} string "Internal Server Error"
// @Router /units [post]
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

// GetUnitHandler godoc
// @Summary Retrieve a specific unit
// @Description Retrieves the details of a unit by its unique ID
// @Tags units
// @Accept json
// @Produce json
// @Param id path int true "Unit ID"
// @Success 200 {object} dto.GetUnitOutput
// @Failure 400 {string} string "Invalid URL or ID"
// @Failure 405 {string} string "Method Not Allowed"
// @Failure 500 {string} string "Internal Server Error"
// @Router /units/{id} [get]
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

// ListUnitHandler godoc
// @Summary Retrieve all units
// @Description Retrieves a list of all units in the system
// @Tags units
// @Accept json
// @Produce json
// @Success 200 {array} dto.GetUnitOutput
// @Failure 405 {string} string "Method Not Allowed"
// @Failure 500 {string} string "Internal Server Error"
// @Router /units [get]
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

// UpdateUnitHandler godoc
// @Summary Update a unit
// @Description Updates the details of a specific unit by its ID
// @Tags units
// @Accept json
// @Produce json
// @Param id path int true "Unit ID"
// @Param unit body dto.CreateUnitInput true "Updated unit data"
// @Success 200 {object} dto.GetUnitOutput
// @Failure 400 {string} string "Invalid URL, ID, or input data"
// @Failure 405 {string} string "Method Not Allowed"
// @Failure 500 {string} string "Internal Server Error"
// @Router /units/{id} [put]
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

// DeleteUnitHandler godoc
// @Summary Delete a unit
// @Description Deletes a unit by its ID
// @Tags units
// @Accept json
// @Produce json
// @Param id path int true "Unit ID"
// @Success 200 {object} map[string]string "Unit deleted successfully"
// @Failure 400 {string} string "Invalid URL or ID"
// @Failure 405 {string} string "Method Not Allowed"
// @Failure 500 {string} string "Internal Server Error"
// @Router /units/{id} [delete]
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

// CreatePaymentHandler godoc
// @Summary Create a new payment
// @Description Adds a new payment to the system
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body dto.CreatePaymentInput true "Payment data"
// @Success 201 {object} dto.CreatePaymentOutput
// @Failure 400 {string} string "Invalid input"
// @Failure 405 {string} string "Method Not Allowed"
// @Failure 500 {string} string "Internal Server Error"
// @Router /payments [post]
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

// ListPaymentsByUnitHandler godoc
// @Summary List payments for a specific unit
// @Description Retrieves a list of all payments associated with a specific unit ID
// @Tags payments
// @Accept json
// @Produce json
// @Param unit_id path int true "Unit ID"
// @Success 200 {array} dto.CreatePaymentOutput
// @Failure 400 {string} string "Invalid URL or ID"
// @Failure 405 {string} string "Method Not Allowed"
// @Failure 500 {string} string "Internal Server Error"
// @Router /payments/unit/{unit_id} [get]
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

// DeletePaymentHandler godoc
// @Summary Delete a payment
// @Description Deletes a payment by its ID
// @Tags payments
// @Accept json
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {object} map[string]string "Payment deleted successfully"
// @Failure 400 {string} string "Invalid URL or ID"
// @Failure 405 {string} string "Method Not Allowed"
// @Failure 500 {string} string "Internal Server Error"
// @Router /payments/{id} [delete]
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

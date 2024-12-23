package buildings

import (
	"billing-sys/internal/domain"
	"billing-sys/internal/domain/services"
)

type CalculateBuildingChargeUseCase struct {
	UnitRepo         domain.UnitRepository
	BuildingRepo     domain.BuildingRepository
	ChargeCalculator *services.ChargeCalculator
}

func (uc *CalculateBuildingChargeUseCase) Execute(buildingID uint) (map[uint]float64, error) {
	building, err := uc.BuildingRepo.GetByID(buildingID)
	if err != nil {
		return nil, err
	}

	units, err := uc.BuildingRepo.GetByBuildingID(buildingID)
	if err != nil {
		return nil, err
	}

	// total charge for each unit
	totalCharges := make(map[uint]float64)

	sharedExpenses := 300.0

	for _, unit := range units {
		charge := uc.ChargeCalculator.Calculate(
			unit.Area,
			building.TotalArea,
			5000.0,
			unit.OccupantsCount,
			sharedExpenses/float64(len(units)),
		)
		totalCharges[unit.ID] = charge
	}

	return totalCharges, nil
}

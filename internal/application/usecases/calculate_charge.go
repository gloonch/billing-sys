package usecases

import (
	"billing-sys/internal/application/dto"
	"billing-sys/internal/domain"
	"errors"
)

type CalculateChargeUseCase struct {
	BuildingRepo domain.BuildingRepository
	UnitRepo     domain.UnitRepository
}

func (uc *CalculateChargeUseCase) Execute(input dto.CalculateChargeInput) (dto.CalculateChargeOutput, error) {

	// Get Building from Repo
	building, err := uc.BuildingRepo.GetByID(input.BuildingID)
	if err != nil {
		return dto.CalculateChargeOutput{}, err
	}
	if building == nil {
		return dto.CalculateChargeOutput{}, errors.New("BuildingID not found")
	}

	// Get Unit from Repo
	unit, err := uc.UnitRepo.GetByID(input.UnitID)
	if err != nil {
		return dto.CalculateChargeOutput{}, err
	}
	if unit == nil {
	}

	// Charge formula
	charge := (unit.Area * 1000) + float64(unit.OccupantCount*5000)

	return dto.CalculateChargeOutput{
		UnitID:     unit.ID,
		BuildingID: building.ID,
		Charge:     charge,
	}, nil
}

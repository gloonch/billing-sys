package units

import (
	"billing-sys/internal/application/dto"
	"billing-sys/internal/domain"
	"errors"
)

type UpdateUnitUseCase struct {
	UnitRepo domain.UnitRepository
}

func (c *UpdateUnitUseCase) Execute(id uint, input dto.CreateUnitInput) (*dto.CreateUnitOutput, error) {

	unit, err := c.UnitRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if unit == nil {
		return nil, errors.New("unit not found")
	}

	unit.Area = input.Area
	unit.UnitNumber = input.UnitNumber
	unit.OccupantsCount = input.OccupantsCount
	unit.BuildingID = input.BuildingID

	err = c.UnitRepo.Save(unit)
	if err != nil {
		return nil, err
	}

	return &dto.CreateUnitOutput{
		ID:             unit.ID,
		Area:           unit.Area,
		UnitNumber:     unit.UnitNumber,
		OccupantsCount: unit.OccupantsCount,
		BuildingID:     unit.BuildingID,
	}, nil
}

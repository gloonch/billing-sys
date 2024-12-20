package units

import (
	"billing-sys/internal/application/dto"
	"billing-sys/internal/domain"
	"billing-sys/internal/domain/entities"
)

type CreateUnitUseCase struct {
	UnitRepo domain.UnitRepository
}

func (uc *CreateUnitUseCase) Execute(input dto.CreateUnitInput) (*dto.CreateUnitOutput, error) {
	unit := &entities.Unit{
		BuildingID:     input.BuildingID,
		UnitNumber:     input.UnitNumber,
		Area:           input.Area,
		OccupantsCount: input.OccupantsCount,
	}

	err := uc.UnitRepo.Save(unit)
	if err != nil {
		return nil, err
	}

	return &dto.CreateUnitOutput{
		ID:             unit.ID,
		UnitNumber:     unit.UnitNumber,
		OccupantsCount: unit.OccupantsCount,
		Area:           unit.Area,
		BuildingID:     unit.BuildingID,
	}, nil
}

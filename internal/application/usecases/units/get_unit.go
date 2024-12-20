package units

import (
	"billing-sys/internal/application/dto"
	"billing-sys/internal/domain"
)

type GetUnitUseCase struct {
	UnitRepo domain.UnitRepository
}

func (uc *GetUnitUseCase) Execute(id uint) (*dto.CreateUnitOutput, error) {

	unit, err := uc.UnitRepo.GetByID(id)
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

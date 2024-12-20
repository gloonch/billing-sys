package units

import (
	"billing-sys/internal/application/dto"
	"billing-sys/internal/domain"
)

type ListAllUnitUseCase struct {
	UnitRepo domain.UnitRepository
}

func (uc *ListAllUnitUseCase) Execute() ([]dto.CreateUnitOutput, error) {

	units, err := uc.UnitRepo.GetAll()

	if err != nil {
		return nil, err
	}

	var result []dto.CreateUnitOutput
	for _, u := range units {
		result = append(result, dto.CreateUnitOutput{
			ID:             u.ID,
			UnitNumber:     u.UnitNumber,
			OccupantsCount: u.OccupantsCount,
			Area:           u.Area,
			BuildingID:     u.BuildingID,
		})
	}

	return result, nil
}

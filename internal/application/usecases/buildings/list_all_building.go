package buildings

import (
	"billing-sys/internal/application/dto"
	"billing-sys/internal/domain"
)

type ListAllBuildingUseCase struct {
	BuildingRepo domain.BuildingRepository
}

func (c *ListAllBuildingUseCase) Execute() ([]dto.GetBuildingOutput, error) {

	buildings, err := c.BuildingRepo.GetAll()

	if err != nil {
		return nil, err
	}

	var result []dto.GetBuildingOutput
	for _, b := range buildings {
		result = append(result, dto.GetBuildingOutput{
			ID:         b.ID,
			Name:       b.Name,
			Address:    b.Address,
			TotalUnits: b.TotalUnits,
			TotalArea:  b.TotalArea,
		})
	}

	return result, nil
}

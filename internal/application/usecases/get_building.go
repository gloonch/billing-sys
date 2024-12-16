package usecases

import (
	"billing-sys/internal/application/dto"
	"billing-sys/internal/domain"
	"errors"
)

type GetBuildingUseCase struct {
	BuildingRepo domain.BuildingRepository
}

func (c *GetBuildingUseCase) Execute(id uint) (*dto.GetBuildingOutput, error) {

	building, err := c.BuildingRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("building not found")
	}

	return &dto.GetBuildingOutput{
		ID:         building.ID,
		Name:       building.Name,
		Address:    building.Address,
		TotalUnits: building.TotalUnits,
		TotalArea:  building.TotalArea,
	}, nil
}

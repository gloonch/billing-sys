package usecases

import (
	"billing-sys/internal/application/dto"
	"billing-sys/internal/domain"
	"billing-sys/internal/domain/entities"
)

type CreateBuildingUseCase struct {
	BuildingRepo domain.BuildingRepository
}

func (c *CreateBuildingUseCase) Execute(input dto.CreateBuildingInput) (*dto.CreateBuildingOutput, error) {
	building := &entities.Building{
		Name:       input.Name,
		Address:    input.Address,
		TotalUnits: input.TotalUnits,
		TotalArea:  input.TotalArea,
	}

	err := c.BuildingRepo.Save(building)
	if err != nil {
		return nil, err
	}

	return &dto.CreateBuildingOutput{
		ID:         building.ID,
		Name:       building.Name,
		Address:    building.Address,
		TotalUnits: building.TotalUnits,
		TotalArea:  building.TotalArea,
	}, nil
}

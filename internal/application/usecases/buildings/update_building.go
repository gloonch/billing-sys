package buildings

import (
	"billing-sys/internal/application/dto"
	"billing-sys/internal/domain"
	"errors"
)

type UpdateBuildingUseCase struct {
	BuildingRepo domain.BuildingRepository
}

func (c *UpdateBuildingUseCase) Execute(id uint, input dto.UpdateBuildingInput) (*dto.UpdateBuildingOutput, error) {

	building, err := c.BuildingRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if building == nil {
		return nil, errors.New("building not found")
	}

	building.Name = input.Name
	building.Address = input.Address
	building.TotalUnits = input.TotalUnits
	building.TotalArea = input.TotalArea

	err = c.BuildingRepo.Save(building)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateBuildingOutput{
		ID:         building.ID,
		Name:       building.Name,
		Address:    building.Address,
		TotalUnits: building.TotalUnits,
		TotalArea:  building.TotalArea,
	}, nil
}

package usecases

import (
	"billing-sys/internal/domain"
)

type DeleteBuildingUseCase struct {
	BuildingRepo domain.BuildingRepository
}

func (c *DeleteBuildingUseCase) Execute(id uint) error {
	err := c.BuildingRepo.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}

package units

import (
	"billing-sys/internal/domain"
)

type DeleteUnitUseCase struct {
	UnitRepo domain.UnitRepository
}

func (uc *DeleteUnitUseCase) Execute(id uint) error {

	err := uc.UnitRepo.DeleteByID(id)
	if err != nil {
		return err
	}

	return nil
}

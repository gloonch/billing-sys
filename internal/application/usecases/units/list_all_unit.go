package units

import (
	"billing-sys/internal/application/dto"
	"billing-sys/internal/domain"
)

type ListAllUnitUseCase struct {
	UnitRepo domain.UnitRepository
}

func (uc *ListAllUnitUseCase) Execute() ([]dto.UnitWithPaymentsOutput, error) {

	units, err := uc.UnitRepo.GetAll()

	if err != nil {
		return nil, err
	}

	var result []dto.UnitWithPaymentsOutput
	for _, unit := range units {
		payments := []dto.PaymentOutput{}
		for _, payment := range unit.Payments {
			payments = append(payments, dto.PaymentOutput{
				ID:          payment.ID,
				Amount:      payment.Amount,
				PaymentDate: payment.PaymentDate,
				Description: payment.Description,
			})
		}
		result = append(result, dto.UnitWithPaymentsOutput{
			ID:             unit.ID,
			BuildingID:     unit.BuildingID,
			UnitNumber:     unit.UnitNumber,
			Area:           unit.Area,
			OccupantsCount: unit.OccupantsCount,
			Payments:       payments,
		})
	}

	return result, nil
}

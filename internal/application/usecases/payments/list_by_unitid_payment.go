package payments

import (
	"billing-sys/internal/application/dto"
	"billing-sys/internal/domain"
)

type ListPaymentsByUnitIDUseCase struct {
	PaymentRepo domain.PaymentRepository
}

func (uc *ListPaymentsByUnitIDUseCase) Execute(unitID uint) ([]dto.PaymentOutput, error) {
	payments, err := uc.PaymentRepo.GetByUnitID(unitID)
	if err != nil {
		return nil, err
	}

	var results []dto.PaymentOutput
	for _, payment := range payments {
		results = append(results, dto.PaymentOutput{
			ID:          payment.ID,
			UnitID:      payment.UnitID,
			Amount:      payment.Amount,
			PaymentDate: payment.PaymentDate,
			Description: payment.Description,
		})
	}
	return results, nil
}

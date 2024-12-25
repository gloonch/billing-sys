package payments

import (
	"billing-sys/internal/application/dto"
	"billing-sys/internal/domain"
	"billing-sys/internal/domain/entities"
)

type CreatePaymentUseCase struct {
	PaymentRepo domain.PaymentRepository
}

func (uc *CreatePaymentUseCase) Execute(input dto.CreatePaymentInput) (*dto.CreatePaymentOutput, error) {
	payment := &entities.Payment{
		UnitID:      input.UnitID,
		Amount:      input.Amount,
		PaymentDate: input.PaymentDate,
		Description: input.Description,
	}

	err := uc.PaymentRepo.Save(payment)
	if err != nil {
		return nil, err
	}

	return &dto.CreatePaymentOutput{
		ID:          payment.ID,
		UnitID:      payment.UnitID,
		Amount:      payment.Amount,
		PaymentDate: payment.PaymentDate,
		Description: payment.Description,
	}, nil
}

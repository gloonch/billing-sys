package payments

import (
	"billing-sys/internal/domain"
)

type DeletePaymentUseCase struct {
	PaymentRepo domain.PaymentRepository
}

func (uc *DeletePaymentUseCase) Execute(paymentID uint) error {
	return uc.PaymentRepo.DeleteByID(paymentID)
}

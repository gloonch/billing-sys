package services

import "billing-sys/internal/domain/entities"

type BillingService struct {
}

func (s *BillingService) CalculateCharge(b entities.Buildings, u entities.Unit) (float64, error) {
	return 0, nil
}

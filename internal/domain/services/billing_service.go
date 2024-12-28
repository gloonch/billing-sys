package services

import "billing-sys/internal/domain/strategies"

type ChargeCalculator struct {
	Strategy strategies.ChargeCalculationStrategy
}

func (c *ChargeCalculator) Calculate(unitArea, totalArea, baseCharge float64, occupantsCount int, sharedExpenses float64) float64 {
	return c.Strategy.Calculate(unitArea, totalArea, baseCharge, occupantsCount, sharedExpenses)
}

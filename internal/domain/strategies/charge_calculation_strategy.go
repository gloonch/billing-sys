package strategies

type ChargeCalculationStrategy interface {
	Calculate(unitArea, totalArea, baseCharge float64, occupantsCount int, sharedExpenses float64) float64
}

type AreaBasedStrategy struct{}

func (s *AreaBasedStrategy) Calculate(unitArea, totalArea, baseCharge float64, occupantsCount int, sharedExpenses float64) float64 {
	return (unitArea / totalArea) * baseCharge
}

type OccupantBasedStrategy struct{}

func (s *OccupantBasedStrategy) Calculate(unitArea, totalArea, baseCharge float64, occupantsCount int, sharedExpenses float64) float64 {
	return float64(occupantsCount) * sharedExpenses
}

type CombinedStrategy struct{}

func (s *CombinedStrategy) Calculate(unitArea, totalArea, baseCharge float64, occupantsCount int, sharedExpenses float64) float64 {
	areaCharge := (unitArea / totalArea) * baseCharge
	occupantCharge := float64(occupantsCount) * sharedExpenses
	return areaCharge + occupantCharge
}

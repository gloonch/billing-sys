package services

type ChargeCalculator struct {
	PerOccupantCharge float64
}

func (c *ChargeCalculator) Calculate(unitArea, totalArea, baseCharge float64, occupantsCount int, sharedExpenses float64) float64 {
	unitCharge := (unitArea / totalArea) * baseCharge
	extraCharge := float64(occupantsCount) * c.PerOccupantCharge
	return unitCharge + extraCharge + sharedExpenses
}

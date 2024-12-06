package dto

type CalculateChargeInput struct {
	BuildingID uint
	UnitID     uint
}

type CalculateChargeOutput struct {
	BuildingID uint
	UnitID     uint
	Charge     float64
}

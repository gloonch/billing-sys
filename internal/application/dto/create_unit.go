package dto

type CreateUnitInput struct {
	UnitNumber     string
	OccupantsCount int
	Area           float64
	BuildingID     uint
}

type CreateUnitOutput struct {
	ID             uint
	UnitNumber     string
	OccupantsCount int
	Area           float64
	BuildingID     uint
}

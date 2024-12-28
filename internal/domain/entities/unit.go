package entities

type Unit struct {
	ID             uint
	BuildingID     uint
	UnitNumber     string
	Area           float64
	OccupantsCount int
	//Residents     []Resident
	Payments []Payment
}

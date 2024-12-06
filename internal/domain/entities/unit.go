package entities

type Unit struct {
	ID            uint
	BuildingID    uint
	UnitNumber    string
	Floor         int
	Area          float64
	OccupantCount int
	Residents     []Resident
	Payments      []Payment
}

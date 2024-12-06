package entities

type Building struct {
	ID         uint
	Name       string
	Address    string
	TotalUnits int
	TotalArea  float64
	Units      []Unit
}

func (b Building) AddUnit(u Unit) {
	b.Units = append(b.Units, u)
	b.TotalUnits = len(b.Units)
}

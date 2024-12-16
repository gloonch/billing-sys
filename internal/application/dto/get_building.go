package dto

type GetBuildingOutput struct {
	ID         uint
	Name       string
	Address    string
	TotalUnits int
	TotalArea  float64
}

package dto

type UpdateBuildingInput struct {
	Name       string
	Address    string
	TotalUnits int
	TotalArea  float64
}

type UpdateBuildingOutput struct {
	ID         uint
	Name       string
	Address    string
	TotalUnits int
	TotalArea  float64
}

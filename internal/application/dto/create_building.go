package dto

type CreateBuildingInput struct {
	Name       string
	Address    string
	TotalUnits int
	TotalArea  float64
}

type CreateBuildingOutput struct {
	ID         uint
	Name       string
	Address    string
	TotalUnits int
	TotalArea  float64
}

package dto

type CreateBuildingInput struct {
	Name       string  `json:"name" example:"Building 1"`
	Address    string  `json:"address" example:"123 Main St"`
	TotalUnits int     `json:"total_units" example:"1"`
	TotalArea  float64 `json:"total_area" example:"10.0"`
}

type CreateBuildingOutput struct {
	ID         uint    `json:"id" example:"1"`
	Name       string  `json:"name" example:"Building 1"`
	Address    string  `json:"address" example:"123 Main St"`
	TotalUnits int     `json:"total_units" example:"1"`
	TotalArea  float64 `json:"total_area" example:"10.0"`
}

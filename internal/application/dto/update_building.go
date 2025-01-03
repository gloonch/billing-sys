package dto

type UpdateBuildingInput struct {
	ID         uint    `json:"id" binding:"required" example:"1"`
	Name       string  `json:"name" binding:"required" example:"Building"`
	Address    string  `json:"address" binding:"required" example:"1st Street"`
	TotalUnits int     `json:"total_units" binding:"required" example:"22"`
	TotalArea  float64 `json:"total_area" binding:"required" example:"20.20"`
}

type UpdateBuildingOutput struct {
	ID         uint    `json:"id" binding:"required" example:"1"`
	Name       string  `json:"name" binding:"required" example:"Building"`
	Address    string  `json:"address" binding:"required" example:"1st Street"`
	TotalUnits int     `json:"total_units" binding:"required" example:"22"`
	TotalArea  float64 `json:"total_area" binding:"required" example:"20.20"`
}

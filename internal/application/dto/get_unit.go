package dto

type GetUnitOutput struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	TotalUnits int     `json:"total_units"`
	TotalArea  float64 `json:"total_area"`
}

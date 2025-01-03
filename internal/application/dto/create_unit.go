package dto

type CreateUnitInput struct {
	UnitNumber     string  `json:"unit_number" example:"100"`
	OccupantsCount int     `json:"occupants_count" example:"10"`
	Area           float64 `json:"area" example:"100"`
	BuildingID     uint    `json:"building_id" example:"0"`
}

type CreateUnitOutput struct {
	ID             uint    `json:"id"`
	UnitNumber     string  `json:"unit_number"`
	OccupantsCount int     `json:"occupants_count"`
	Area           float64 `json:"area"`
	BuildingID     uint    `json:"building_id"`
}

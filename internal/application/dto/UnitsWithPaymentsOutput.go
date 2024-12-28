package dto

type UnitWithPaymentsOutput struct {
	ID             uint            `json:"id"`
	BuildingID     uint            `json:"building_id"`
	UnitNumber     string          `json:"unit_number"`
	Area           float64         `json:"area"`
	OccupantsCount int             `json:"occupants_count"`
	Payments       []PaymentOutput `json:"payments"`
}

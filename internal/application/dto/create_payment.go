package dto

import "time"

type CreatePaymentInput struct {
	UnitID      uint
	Amount      float64
	PaymentDate time.Time
	Description string
}

type CreatePaymentOutput struct {
	ID          uint
	UnitID      uint
	Amount      float64
	PaymentDate time.Time
	Description string
}

type PaymentOutput struct {
	ID          uint      `json:"id"`
	UnitID      uint      `json:"unit_id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Description string    `json:"description"`
}

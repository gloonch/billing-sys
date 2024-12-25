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
	ID          uint
	UnitID      uint
	Amount      float64
	PaymentDate time.Time
	Description string
}

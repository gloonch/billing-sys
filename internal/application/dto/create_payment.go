package dto

import "time"

type CreatePaymentInput struct {
	UnitID      uint      `json:"unit_id" example:"1" binding:"required"`
	Amount      float64   `json:"amount" example:"10.0" binding:"required"`
	PaymentDate time.Time `json:"payment_date" example:"2020-01-01" binding:"required"`
	Description string    `json:"description" example:"some description here" binding:"required"`
}

type CreatePaymentOutput struct {
	ID          uint
	UnitID      uint      `json:"unit_id" example:"1" binding:"required"`
	Amount      float64   `json:"amount" example:"10.0" binding:"required"`
	PaymentDate time.Time `json:"payment_date" example:"2020-01-01" binding:"required"`
	Description string    `json:"description" example:"some description here" binding:"required"`
}

type PaymentOutput struct {
	ID          uint      `json:"id"`
	UnitID      uint      `json:"unit_id" example:"1" binding:"required"`
	Amount      float64   `json:"amount" example:"10.0" binding:"required"`
	PaymentDate time.Time `json:"payment_date" example:"2020-01-01" binding:"required"`
	Description string    `json:"description" example:"some description here" binding:"required"`
}

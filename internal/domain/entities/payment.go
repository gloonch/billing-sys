package entities

import "time"

type Payment struct {
	ID          uint
	UnitID      uint
	Amount      float64
	PaymentDate time.Time
	Description string
}

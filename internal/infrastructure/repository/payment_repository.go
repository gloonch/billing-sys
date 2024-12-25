package repository

import (
	"billing-sys/internal/domain"
	"billing-sys/internal/domain/entities"
	"database/sql"
	"fmt"
)

type PgPaymentRepository struct {
	db *sql.DB
}

func NewPgPaymentRepository(db *sql.DB) domain.PaymentRepository {
	return &PgPaymentRepository{db: db}
}

func (r *PgPaymentRepository) Save(payment *entities.Payment) error {
	// TODO: create a validation class
	if payment.UnitID == 0 {
		return fmt.Errorf("validation error: UnitID is required")
	}

	if payment.Amount <= 0 {
		return fmt.Errorf("validation error: Amount must be greater than 0")
	}

	if payment.PaymentDate.IsZero() {
		return fmt.Errorf("validation error: PaymentDate is required")
	}

	var exists bool
	queryCheck := `SELECT EXISTS (SELECT 1 FROM units WHERE id = $1)`
	err := r.db.QueryRow(queryCheck, payment.UnitID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("database error: failed to check UnitID existence: %v", err)
	}
	if !exists {
		return fmt.Errorf("validation error: UnitID %d does not exist", payment.UnitID)
	}

	queryInsert := `
        INSERT INTO payments (unit_id, amount, payment_date, description)
        VALUES ($1, $2, $3, $4) RETURNING id
    `
	err = r.db.QueryRow(queryInsert, payment.UnitID, payment.Amount, payment.PaymentDate, payment.Description).Scan(&payment.ID)
	if err != nil {
		return fmt.Errorf("database error: failed to save payment: %v", err)
	}

	return nil
}

func (r *PgPaymentRepository) GetByID(id uint) (*entities.Payment, error) {
	query := `SELECT id, unit_id, amount, payment_date, description FROM payments WHERE id = $1`
	row := r.db.QueryRow(query, id)

	payment := &entities.Payment{}
	err := row.Scan(&payment.ID, &payment.UnitID, &payment.Amount, &payment.PaymentDate, &payment.Description)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *PgPaymentRepository) GetByUnitID(unitID uint) ([]entities.Payment, error) {
	query := `SELECT id, unit_id, amount, payment_date, description FROM payments WHERE unit_id = $1`
	rows, err := r.db.Query(query, unitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []entities.Payment
	for rows.Next() {
		payment := entities.Payment{}
		err := rows.Scan(&payment.ID, &payment.UnitID, &payment.Amount, &payment.PaymentDate, &payment.Description)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func (r *PgPaymentRepository) DeleteByID(id uint) error {
	query := `DELETE FROM payments WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

package repository

import (
	"billing-sys/internal/domain"
	"billing-sys/internal/domain/entities"
	"database/sql"
	"errors"
)

type PgUnitRepository struct {
	db *sql.DB
}

func NewPgUnitRepository(db *sql.DB) domain.UnitRepository {
	return &PgUnitRepository{db: db}
}

func (r *PgUnitRepository) GetAll() ([]entities.Unit, error) {
	query := `
        SELECT 
            u.id AS unit_id, 
            u.building_id, 
            u.unit_number, 
            u.area, 
            u.occupants_count, 
            p.id AS payment_id, 
            p.amount, 
            p.payment_date, 
            p.description
        FROM units u
        LEFT JOIN payments p ON u.id = p.unit_id
    `
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	unitMap := make(map[uint]*entities.Unit)

	for rows.Next() {
		var (
			unitID      uint
			paymentID   sql.NullInt64
			amount      sql.NullFloat64
			paymentDate sql.NullTime
			description sql.NullString
			unit        entities.Unit
		)

		err := rows.Scan(
			&unitID, &unit.BuildingID, &unit.UnitNumber, &unit.Area, &unit.OccupantsCount,
			&paymentID, &amount, &paymentDate, &description,
		)
		if err != nil {
			return nil, err
		}

		// بررسی اینکه آیا این واحد قبلاً اضافه شده است یا خیر
		if _, exists := unitMap[unitID]; !exists {
			unit.ID = unitID
			unit.Payments = []entities.Payment{}
			unitMap[unitID] = &unit
		}

		// اگر پرداخت وجود دارد، آن را اضافه کنید
		if paymentID.Valid {
			unitMap[unitID].Payments = append(unitMap[unitID].Payments, entities.Payment{
				ID:          uint(paymentID.Int64),
				UnitID:      unitID,
				Amount:      amount.Float64,
				PaymentDate: paymentDate.Time,
				Description: description.String,
			})
		}
	}

	var units []entities.Unit
	for _, unit := range unitMap {
		units = append(units, *unit)
	}

	return units, nil
}

func (r *PgUnitRepository) GetByID(id uint) (*entities.Unit, error) {
	var u entities.Unit

	row := r.db.QueryRow(`SELECT id, unit_number, area, occupants_count FROM units WHERE id = $1`, id)
	err := row.Scan(&u.ID, &u.UnitNumber, &u.Area, &u.OccupantsCount)
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		// there was no response and no error
		return nil, nil
	}

	return &u, nil
}

func (r *PgUnitRepository) Save(u *entities.Unit) error {

	// if u.ID == 0 it is an insert
	if u.ID == 0 {
		err := r.db.QueryRow("INSERT INTO units (unit_number, area, occupants_count, building_id) VALUES ($1, $2, $3, $4) RETURNING id",
			u.UnitNumber, u.Area, u.OccupantsCount, u.BuildingID).Scan(&u.ID)
		return err
	} else {
		_, err := r.db.Exec("UPDATE units SET unit_number=$1, area=$2, occupants_count=$3 WHERE id=$4",
			u.UnitNumber, u.Area, u.OccupantsCount, u.ID)
		return err
	}
}

func (r *PgUnitRepository) DeleteByID(id uint) error {
	result, err := r.db.Exec("DELETE FROM units WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no unit found with given id")
	}

	return nil
}

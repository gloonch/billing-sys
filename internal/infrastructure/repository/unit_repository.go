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
	rows, err := r.db.Query("SELECT * FROM unit")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var units []entities.Unit
	for rows.Next() {
		var u entities.Unit
		if err := rows.Scan(&u.ID, &u.BuildingID, &u.UnitNumber, &u.Floor, &u.Area, &u.OccupantCount); err != nil {
			return nil, err
		}
		units = append(units, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return units, nil
}

func (r *PgUnitRepository) GetByID(id uint) (*entities.Unit, error) {
	var u entities.Unit

	// TODO: fix this to
	row := r.db.QueryRow(`SELECT id, unit_number, floor, area, occupant_count FROM units WHERE id = $1`, id)
	err := row.Scan(&u.ID, &u.UnitNumber, &u.Floor, &u.Area, &u.OccupantCount)
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
		err := r.db.QueryRow("INSERT INTO units (unit_number, floor, area, occupant_count) VALUES ($1, $2, $3, $4) RETURNING id",
			u.UnitNumber, u.Floor, u.Area, u.OccupantCount).Scan(&u.ID)
		return err
	} else {
		_, err := r.db.Exec("UPDATE units SET unit_number=$1, unit_number=$2, unit_number=$3, unit_number=$4, ",
			u.UnitNumber, u.Floor, u.Area, u.OccupantCount, u.ID)
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

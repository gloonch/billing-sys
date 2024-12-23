package repository

import (
	"billing-sys/internal/domain"
	"billing-sys/internal/domain/entities"
	"database/sql"
	"errors"
)

type PgBuildingRepository struct {
	db *sql.DB
}

func NewPgBuildingRepository(db *sql.DB) domain.BuildingRepository {
	return &PgBuildingRepository{db: db}
}

func (r *PgBuildingRepository) GetAll() ([]entities.Buildings, error) {
	rows, err := r.db.Query("SELECT id, name, address, total_units, total_area FROM buildings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var buildings []entities.Buildings
	for rows.Next() {
		var b entities.Buildings
		if err := rows.Scan(&b.ID, &b.Name, &b.Address, &b.TotalUnits, &b.TotalArea); err != nil {
			return nil, err
		}
		buildings = append(buildings, b)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return buildings, nil
}

func (r *PgBuildingRepository) GetByID(id uint) (*entities.Buildings, error) {
	var b entities.Buildings

	row := r.db.QueryRow("SELECT id, name, address, total_units, total_area FROM buildings WHERE id = $1", id)
	err := row.Scan(&b.ID, &b.Name, &b.Address, &b.TotalUnits, &b.TotalArea)
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		// no response, no error
		return nil, nil
	}

	return &b, nil
}

func (r *PgBuildingRepository) Save(b *entities.Buildings) error {

	// if b.ID == 0 it is an insert
	if b.ID == 0 {
		err := r.db.QueryRow(
			"INSERT INTO buildings (name, address, total_units, total_area) VALUES ($1, $2, $3, $4) RETURNING id",
			b.Name, b.Address, b.TotalUnits, b.TotalArea).Scan(&b.ID)
		return err
	} else {
		_, err := r.db.Exec(
			"UPDATE buildings SET name=$1, address=$2, total_units=$3, total_area=$4 WHERE id=$5 ",
			b.Name, b.Address, b.TotalUnits, b.TotalArea, b.ID)
		return err
	}
}

func (r *PgBuildingRepository) DeleteByID(id uint) error {
	result, err := r.db.Exec("DELETE FROM buildings WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no building found with given id")
	}

	return nil
}

func (r *PgBuildingRepository) GetByBuildingID(buildingID uint) ([]entities.Unit, error) {
	// fetch units by building ID
	query := `SELECT id, unit_number, occupants_count, area, building_id FROM units WHERE building_id = $1`

	rows, err := r.db.Query(query, buildingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize slice to store units
	var units []entities.Unit

	// Iterate through the rows
	for rows.Next() {
		var unit entities.Unit
		err := rows.Scan(&unit.ID, &unit.UnitNumber, &unit.OccupantsCount, &unit.Area, &unit.BuildingID)
		if err != nil {
			return nil, err
		}
		units = append(units, unit)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return units, nil
}

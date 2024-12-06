package domain

import "billing-sys/internal/domain/entities"

type BuildingRepository interface {
	GetByID(id uint) (*entities.Building, error)
	Save(building *entities.Building) error
}

type UnitRepository interface {
	GetByID(id uint) (*entities.Unit, error)
	Save(unit *entities.Unit) error
}

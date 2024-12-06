package domain

import "billing-sys/internal/domain/entities"

type BuildingRepository interface {
	GetByID(id uint) (*entities.Building, error)
	Save(building *entities.Building) error
	GetAll() ([]entities.Building, error)
	DeleteByID(id uint) error
}

type UnitRepository interface {
	GetByID(id uint) (*entities.Unit, error)
	Save(unit *entities.Unit) error
	GetAll() ([]entities.Unit, error)
	DeleteByID(id uint) error
}

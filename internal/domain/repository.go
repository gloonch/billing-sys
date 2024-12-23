package domain

import "billing-sys/internal/domain/entities"

type BuildingRepository interface {
	GetByID(id uint) (*entities.Buildings, error)
	Save(building *entities.Buildings) error
	GetAll() ([]entities.Buildings, error)
	DeleteByID(id uint) error
	GetByBuildingID(buildingID uint) ([]entities.Unit, error)
}

type UnitRepository interface {
	GetByID(id uint) (*entities.Unit, error)
	Save(unit *entities.Unit) error
	GetAll() ([]entities.Unit, error)
	DeleteByID(id uint) error
}

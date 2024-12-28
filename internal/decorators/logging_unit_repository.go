package decorators

import (
	"billing-sys/internal/domain"
	"billing-sys/internal/domain/entities"
	"log"
)

type LoggingUnitRepository struct {
	Repo domain.UnitRepository
}

func (l *LoggingUnitRepository) GetByID(id uint) (*entities.Unit, error) {
	log.Printf("GetByID called with id: %d", id)
	unit, err := l.Repo.GetByID(id)
	if err != nil {
		log.Printf("GetByID failed with error: %v", err)
		return nil, err
	}
	log.Printf("GetByID successful: %+v", unit)
	return unit, nil
}

func (l *LoggingUnitRepository) Save(unit *entities.Unit) error {
	log.Printf("Save called with unit: %+v", unit)
	err := l.Repo.Save(unit)
	if err != nil {
		log.Printf("Save failed with error: %v", err)
		return err
	}
	log.Printf("Save successful")
	return nil
}

func (l *LoggingUnitRepository) GetAll() ([]entities.Unit, error) {
	log.Println("GetAll called")
	units, err := l.Repo.GetAll()
	if err != nil {
		log.Printf("GetAll failed with error: %v", err)
		return nil, err
	}
	log.Printf("GetAll successful: %d units fetched", len(units))
	return units, nil
}

func (l *LoggingUnitRepository) DeleteByID(id uint) error {
	log.Printf("DeleteByID called with id: %d", id)
	err := l.Repo.DeleteByID(id)
	if err != nil {
		log.Printf("DeleteByID failed with error: %v", err)
		return err
	}
	log.Printf("DeleteByID successful")
	return nil
}

package decorators

import (
	"billing-sys/internal/domain"
	"billing-sys/internal/domain/entities"
	"log"
)

type LoggingBuildingRepository struct {
	Repo domain.BuildingRepository
}

func (l *LoggingBuildingRepository) GetByID(id uint) (*entities.Buildings, error) {
	log.Printf("GetByID called with id: %d", id)
	building, err := l.Repo.GetByID(id)
	if err != nil {
		log.Printf("GetByID failed with error: %v", err)
		return nil, err
	}
	log.Printf("GetByID successful: %+v", building)
	return building, nil
}

func (l *LoggingBuildingRepository) Save(building *entities.Buildings) error {
	log.Printf("Save called with building: %+v", building)
	err := l.Repo.Save(building)
	if err != nil {
		log.Printf("Save failed with error: %v", err)
		return err
	}
	log.Printf("Save successful")
	return nil
}

func (l *LoggingBuildingRepository) GetAll() ([]entities.Buildings, error) {
	log.Println("GetAll called")
	buildings, err := l.Repo.GetAll()
	if err != nil {
		log.Printf("GetAll failed with error: %v", err)
		return nil, err
	}
	log.Printf("GetAll successful: %d buildings fetched", len(buildings))
	return buildings, nil
}

func (l *LoggingBuildingRepository) DeleteByID(id uint) error {
	log.Printf("DeleteByID called with id: %d", id)
	err := l.Repo.DeleteByID(id)
	if err != nil {
		log.Printf("DeleteByID failed with error: %v", err)
		return err
	}
	log.Printf("DeleteByID successful")
	return nil
}

func (l *LoggingBuildingRepository) GetByBuildingID(buildingID uint) ([]entities.Unit, error) {
	log.Printf("GetByBuildingID called with buildingID: %d", buildingID)
	units, err := l.Repo.GetByBuildingID(buildingID)
	if err != nil {
		log.Printf("GetByBuildingID failed with error: %v", err)
		return nil, err
	}
	log.Printf("GetByBuildingID successful: %d units fetched", len(units))
	return units, nil
}

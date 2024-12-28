package decorators

import (
	"billing-sys/internal/domain"
	"billing-sys/internal/domain/entities"
	"billing-sys/internal/utils"
	"strconv"
)

type LoggingBuildingRepository struct {
	Repo domain.BuildingRepository
}

func (l *LoggingBuildingRepository) GetByID(id uint) (*entities.Buildings, error) {
	utils.LogInfo("Building", "GetByID called with id: "+strconv.Itoa(int(id)), "GetByID(id uint) (*entities.Buildings, error)")
	building, err := l.Repo.GetByID(id)
	if err != nil {
		utils.LogError("Building", "GetByID failed with error: "+err.Error(), "GetByID(id uint) (*entities.Buildings, error)")
		return nil, err
	}
	utils.LogSuccess("Building", "GetByID successful", "GetByID(id uint) (*entities.Buildings, error)")
	return building, nil
}

func (l *LoggingBuildingRepository) Save(building *entities.Buildings) error {
	utils.LogInfo("Building", "Save called", "Save(building *entities.Buildings) error")
	err := l.Repo.Save(building)
	if err != nil {
		utils.LogError("Building", "Save failed with error: "+err.Error(), "Save(building *entities.Buildings) error")
		return err
	}
	utils.LogInfo("Building", "Save successful", "Save(building *entities.Buildings) error")
	return nil
}

func (l *LoggingBuildingRepository) GetAll() ([]entities.Buildings, error) {
	utils.LogInfo("Building", "GetAll called", "GetAll() ([]entities.Buildings, error)")
	buildings, err := l.Repo.GetAll()
	if err != nil {
		utils.LogError("Building", "GetAll failed with error: "+err.Error(), "GetAll() ([]entities.Buildings, error)")
		return nil, err
	}
	utils.LogInfo("Building", "GetAll successful", "GetAll() ([]entities.Buildings, error)")
	return buildings, nil
}

func (l *LoggingBuildingRepository) DeleteByID(id uint) error {
	utils.LogInfo("Building", "DeleteByID called with id: "+strconv.Itoa(int(id)), "DeleteByID(id uint) error")
	err := l.Repo.DeleteByID(id)
	if err != nil {
		utils.LogError("Building", "DeleteByID failed with error: "+err.Error(), "DeleteByID(id uint) error")
		return err
	}
	utils.LogInfo("Building", "DeleteByID successful", "DeleteByID(id uint) error")
	return nil
}

func (l *LoggingBuildingRepository) GetByBuildingID(buildingID uint) ([]entities.Unit, error) {
	utils.LogInfo("Building", "GetByBuildingID called with buildingID: "+strconv.Itoa(int(buildingID)), "GetByBuildingID(buildingID uint) ([]entities.Unit, error)")
	units, err := l.Repo.GetByBuildingID(buildingID)
	if err != nil {
		utils.LogError("Building", "GetByBuildingID failed with error: "+err.Error(), "GetByBuildingID(buildingID uint) ([]entities.Unit, error)")
		return nil, err
	}
	utils.LogInfo("Building", "GetByBuildingID successful", "GetByBuildingID(buildingID uint) ([]entities.Unit, error)")
	return units, nil
}

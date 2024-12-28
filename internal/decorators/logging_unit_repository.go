package decorators

import (
	"billing-sys/internal/domain"
	"billing-sys/internal/domain/entities"
	"billing-sys/internal/utils"
	"strconv"
)

type LoggingUnitRepository struct {
	Repo domain.UnitRepository
}

func (l *LoggingUnitRepository) GetByID(id uint) (*entities.Unit, error) {
	utils.LogInfo("Unit", "GetByID called with id: "+strconv.Itoa(int(id)), "GetByID(id uint) (*entities.Unit, error)")
	unit, err := l.Repo.GetByID(id)
	if err != nil {
		utils.LogError("Unit", "GetByID failed with error: "+err.Error(), "GetByID(id uint) (*entities.Unit, error)")
		return nil, err
	}
	utils.LogSuccess("Unit", "GetByID successful", "GetByID(id uint) (*entities.Unit, error)")
	return unit, nil
}

func (l *LoggingUnitRepository) Save(unit *entities.Unit) error {
	utils.LogInfo("Unit", "Save called with unit: ", "Save(unit *entities.Unit) error")
	err := l.Repo.Save(unit)
	if err != nil {
		utils.LogError("Unit", "Save failed with error: "+err.Error(), "Save(unit *entities.Unit) error")
		return err
	}
	utils.LogSuccess("Unit", "Save successful", "Save(unit *entities.Unit) error")
	return nil
}

func (l *LoggingUnitRepository) GetAll() ([]entities.Unit, error) {
	utils.LogInfo("Unit", "GetAll called", "GetAll() ([]entities.Unit, error)")
	units, err := l.Repo.GetAll()
	if err != nil {
		utils.LogError("Unit", "GetAll failed with error: "+err.Error(), "GetAll() ([]entities.Unit, error)")
		return nil, err
	}
	utils.LogSuccess("Unit", "GetAll successful"+strconv.Itoa(len(units)), "GetAll() ([]entities.Unit, error)")
	return units, nil
}

func (l *LoggingUnitRepository) DeleteByID(id uint) error {
	utils.LogInfo("Unit", "DeleteByID called with id: "+strconv.Itoa(int(id)), "DeleteByID(id uint) error")
	err := l.Repo.DeleteByID(id)
	if err != nil {
		utils.LogError("Unit", "DeleteByID failed with error: "+err.Error(), "DeleteByID(id uint) error")
		return err
	}
	utils.LogSuccess("Unit", "DeleteByID successful", "DeleteByID(id uint) error")

	return nil
}

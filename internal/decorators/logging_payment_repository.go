package decorators

import (
	"billing-sys/internal/domain"
	"billing-sys/internal/domain/entities"
	"billing-sys/internal/utils"
	"log"
	"strconv"
)

type LoggingPaymentRepository struct {
	Repo domain.PaymentRepository
}

func (l *LoggingPaymentRepository) Save(payment *entities.Payment) error {
	utils.LogInfo("Payment", "Save called with payment", "Save(payment *entities.Payment)")
	err := l.Repo.Save(payment)
	if err != nil {
		utils.LogError("Payment", "Save failed with error: "+err.Error(), "Save(payment *entities.Payment)")
		return err
	}
	utils.LogSuccess("Payment", "Save successful", "Save(payment *entities.Payment)")
	return nil
}

func (l *LoggingPaymentRepository) GetByID(id uint) (*entities.Payment, error) {
	utils.LogInfo("Payment", "GetByID called with id: "+strconv.Itoa(int(id)), "GetByID(id uint) (*entities.Payment, error)")
	payment, err := l.Repo.GetByID(id)
	if err != nil {
		utils.LogError("Payment", "GetByID failed with error: "+err.Error(), "GetByID(id uint) (*entities.Payment, error)")
		log.Printf("GetByID failed with error: %v", err)
		return nil, err
	}
	utils.LogSuccess("Payment", "GetByID successful", "GetByID(id uint) (*entities.Payment, error)")
	return payment, nil
}

func (l *LoggingPaymentRepository) GetByUnitID(unitID uint) ([]entities.Payment, error) {
	utils.LogInfo("Payment", "GetByUnitID called with unitID: "+strconv.Itoa(int(unitID)), "GetByUnitID(unitID uint) ([]entities.Payment, error)")
	payments, err := l.Repo.GetByUnitID(unitID)
	if err != nil {
		utils.LogError("Payment", "GetByUnitID failed with error: "+err.Error(), "GetByUnitID(unitID uint) ([]entities.Payment, error)")
		return nil, err
	}
	utils.LogInfo("Payment", "GetByUnitID successful payments fetched: "+string(len(payments)), "GetByUnitID(unitID uint) ([]entities.Payment, error)")
	return payments, nil
}

func (l *LoggingPaymentRepository) DeleteByID(id uint) error {
	utils.LogInfo("Payment", "DeleteByID called with id: "+strconv.Itoa(int(id)), "DeleteByID(id uint) error")
	err := l.Repo.DeleteByID(id)
	if err != nil {
		utils.LogError("Payment", "DeleteByID failed with error: "+err.Error(), "DeleteByID(id uint) error")
		return err
	}
	utils.LogInfo("Payment", "DeleteByID successful", "DeleteByID(id uint) error")
	return nil
}

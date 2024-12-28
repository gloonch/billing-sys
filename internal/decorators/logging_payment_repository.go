package decorators

import (
	"billing-sys/internal/domain"
	"billing-sys/internal/domain/entities"
	"log"
)

type LoggingPaymentRepository struct {
	Repo domain.PaymentRepository
}

func (l *LoggingPaymentRepository) Save(payment *entities.Payment) error {
	log.Printf("Save called with payment: %+v", payment)
	err := l.Repo.Save(payment)
	if err != nil {
		log.Printf("Save failed with error: %v", err)
		return err
	}
	log.Printf("Save successful")
	return nil
}

func (l *LoggingPaymentRepository) GetByID(id uint) (*entities.Payment, error) {
	log.Printf("GetByID called with id: %d", id)
	payment, err := l.Repo.GetByID(id)
	if err != nil {
		log.Printf("GetByID failed with error: %v", err)
		return nil, err
	}
	log.Printf("GetByID successful: %+v", payment)
	return payment, nil
}

func (l *LoggingPaymentRepository) GetByUnitID(unitID uint) ([]entities.Payment, error) {
	log.Printf("GetByUnitID called with unitID: %d", unitID)
	payments, err := l.Repo.GetByUnitID(unitID)
	if err != nil {
		log.Printf("GetByUnitID failed with error: %v", err)
		return nil, err
	}
	log.Printf("GetByUnitID successful: %d payments fetched", len(payments))
	return payments, nil
}

func (l *LoggingPaymentRepository) DeleteByID(id uint) error {
	log.Printf("DeleteByID called with id: %d", id)
	err := l.Repo.DeleteByID(id)
	if err != nil {
		log.Printf("DeleteByID failed with error: %v", err)
		return err
	}
	log.Printf("DeleteByID successful")
	return nil
}

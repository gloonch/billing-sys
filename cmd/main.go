package main

import (
	_ "billing-sys/docs"
	"billing-sys/internal/application/usecases/buildings"
	"billing-sys/internal/application/usecases/payments"
	"billing-sys/internal/application/usecases/units"
	"billing-sys/internal/decorators"
	"billing-sys/internal/domain/services"
	"billing-sys/internal/infrastructure/database"
	mHttp "billing-sys/internal/infrastructure/http"
	"billing-sys/internal/infrastructure/repository"
	"log"
	"net/http"
)

func main() {

	cfg := database.DBConfig{
		Host:     "postgres",
		Port:     5432,
		Username: "postgres",
		Password: "1234",
		DBName:   "billingsys",
		SSLMode:  "disable",
	}

	// make connection to database
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)

		return
	}
	defer db.Close()
	log.Println("Connected to database")

	// Original Repositories
	buildingRepo := repository.NewPgBuildingRepository(db)
	unitRepo := repository.NewPgUnitRepository(db)
	paymentRepo := repository.NewPgPaymentRepository(db)

	// Logging Decorators
	loggingBuildingRepo := &decorators.LoggingBuildingRepository{Repo: buildingRepo}
	loggingUnitRepo := &decorators.LoggingUnitRepository{Repo: unitRepo}
	loggingPaymentRepo := &decorators.LoggingPaymentRepository{Repo: paymentRepo}

	// Charge Calculator
	chargeCalculator := services.ChargeCalculator{}

	// get buildings use cases
	createBuildingUC := &buildings.CreateBuildingUseCase{
		BuildingRepo: loggingBuildingRepo,
	}
	getBuildingUC := &buildings.GetBuildingUseCase{
		BuildingRepo: loggingBuildingRepo,
	}
	updateBuildingUC := &buildings.UpdateBuildingUseCase{
		BuildingRepo: loggingBuildingRepo,
	}
	listBuildingUC := &buildings.ListAllBuildingUseCase{
		BuildingRepo: loggingBuildingRepo,
	}
	deleteBuildingUC := &buildings.DeleteBuildingUseCase{
		BuildingRepo: loggingBuildingRepo,
	}
	calculateBuildingChargeUC := &buildings.CalculateBuildingChargeUseCase{
		UnitRepo:         loggingUnitRepo,
		BuildingRepo:     loggingBuildingRepo,
		ChargeCalculator: &chargeCalculator,
	}

	// Unit Use Cases
	createUnitUC := &units.CreateUnitUseCase{
		UnitRepo: loggingUnitRepo,
	}
	getUnitUC := &units.GetUnitUseCase{
		UnitRepo: loggingUnitRepo,
	}
	listUnitUC := &units.ListAllUnitUseCase{
		UnitRepo: loggingUnitRepo,
	}
	updateUnitUC := &units.UpdateUnitUseCase{
		UnitRepo: loggingUnitRepo,
	}
	deleteUnitUC := &units.DeleteUnitUseCase{
		UnitRepo: loggingUnitRepo,
	}

	// get unit use cases
	createPaymentUC := &payments.CreatePaymentUseCase{
		PaymentRepo: loggingPaymentRepo,
	}
	deletePaymentUC := &payments.DeletePaymentUseCase{
		PaymentRepo: loggingPaymentRepo,
	}
	listPaymentsByUnitIDUC := &payments.ListPaymentsByUnitIDUseCase{
		PaymentRepo: loggingPaymentRepo,
	}

	// handlers
	handlers := &mHttp.Handlers{
		CreateBuildingUC:          createBuildingUC,
		GetBuildingUC:             getBuildingUC,
		ListBuildingsUC:           listBuildingUC,
		UpdateBuildingUC:          updateBuildingUC,
		DeleteBuildingUC:          deleteBuildingUC,
		CalculateBuildingChargeUC: calculateBuildingChargeUC,

		CreateUnitUC: createUnitUC,
		GetUnitUC:    getUnitUC,
		ListUnitsUC:  listUnitUC,
		UpdateUnitUC: updateUnitUC,
		DeleteUnitUC: deleteUnitUC,

		CreatePaymentUC:        createPaymentUC,
		DeletePaymentUC:        deletePaymentUC,
		ListPaymentsByUnitIDUC: listPaymentsByUnitIDUC,
	}

	// config router
	router := mHttp.NewRouter(handlers)

	port := ":8000"
	log.Println("Starting server on port ", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

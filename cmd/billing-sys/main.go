package main

import (
	"billing-sys/internal/application/usecases/buildings"
	"billing-sys/internal/application/usecases/payments"
	"billing-sys/internal/application/usecases/units"
	"billing-sys/internal/domain/services"
	"billing-sys/internal/infrastructure/database"
	mHttp "billing-sys/internal/infrastructure/http"
	"billing-sys/internal/infrastructure/repository"
	"log"
	"net/http"
)

func main() {

	cfg := database.DBConfig{
		Host:     "localhost",
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

	// get repository
	buildingRepo := repository.NewPgBuildingRepository(db)
	unitRepo := repository.NewPgUnitRepository(db)
	paymentRepo := repository.NewPgPaymentRepository(db)
	chargeCalculator := services.ChargeCalculator{}

	// get buildings use cases
	createBuildingUC := &buildings.CreateBuildingUseCase{
		BuildingRepo: buildingRepo,
	}
	getBuildingUC := &buildings.GetBuildingUseCase{
		BuildingRepo: buildingRepo,
	}
	updateBuildingUC := &buildings.UpdateBuildingUseCase{
		BuildingRepo: buildingRepo,
	}
	listBuildingUC := &buildings.ListAllBuildingUseCase{
		BuildingRepo: buildingRepo,
	}
	deleteBuildingUC := &buildings.DeleteBuildingUseCase{
		BuildingRepo: buildingRepo,
	}
	calculateBuildingChargeUC := &buildings.CalculateBuildingChargeUseCase{
		UnitRepo:         unitRepo,
		BuildingRepo:     buildingRepo,
		ChargeCalculator: &chargeCalculator,
	}

	// get unit use cases
	createUnitUC := &units.CreateUnitUseCase{
		UnitRepo: unitRepo,
	}
	getUnitUC := &units.GetUnitUseCase{
		UnitRepo: unitRepo,
	}
	listUnitUC := &units.ListAllUnitUseCase{
		UnitRepo: unitRepo,
	}
	updateUnitUC := &units.UpdateUnitUseCase{
		UnitRepo: unitRepo,
	}
	deleteUnitUC := &units.DeleteUnitUseCase{
		UnitRepo: unitRepo,
	}

	// get unit use cases
	createPaymentUC := &payments.CreatePaymentUseCase{
		PaymentRepo: paymentRepo,
	}
	deletePaymentUC := &payments.DeletePaymentUseCase{
		PaymentRepo: paymentRepo,
	}
	listPaymentsByUnitIDUC := &payments.ListPaymentsByUnitIDUseCase{
		PaymentRepo: paymentRepo,
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

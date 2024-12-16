package main

import (
	"billing-sys/internal/application/usecases"
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
		Password: "password",
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

	// get use cases
	createBuildingUC := &usecases.CreateBuildingUseCase{
		BuildingRepo: buildingRepo,
	}
	getBuildingUC := &usecases.GetBuildingUseCase{
		BuildingRepo: buildingRepo,
	}
	updateBuildingUC := &usecases.UpdateBuildingUseCase{
		BuildingRepo: buildingRepo,
	}
	listBuildingUC := &usecases.ListAllBuildingUseCase{
		BuildingRepo: buildingRepo,
	}
	deleteBuildingUC := &usecases.DeleteBuildingUseCase{
		BuildingRepo: buildingRepo,
	}

	// handlers
	handlers := &mHttp.Handlers{
		CreateBuildingUC: createBuildingUC,
		GetBuildingUC:    getBuildingUC,
		ListBuildingsUC:  listBuildingUC,
		UpdateBuildingUC: updateBuildingUC,
		DeleteBuildingUC: deleteBuildingUC,
	}

	// config router
	router := mHttp.NewRouter(handlers)

	port := ":8000"
	log.Println("Starting server on port ", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

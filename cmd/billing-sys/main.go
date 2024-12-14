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
	calculateChargeUC := &usecases.CalculateChargeUseCase{
		BuildingRepo: buildingRepo,
		UnitRepo:     nil,
	}

	// handlers
	handlers := &mHttp.Handlers{
		calculateChargeUC,
	}

	// config router
	router := mHttp.NewRouter(handlers)

	port := ":8000"
	log.Println("Starting server on port ", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

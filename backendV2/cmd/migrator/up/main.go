package main

import (
	"log"
	"queueAppV2/internal/config"
	"queueAppV2/internal/repository/postgres"
	"queueAppV2/internal/repository/postgres/migrations"
)

func main() {
	cfg := config.MustLoad()
	db := postgres.MustCreate(cfg)
	migrations.ApplyMigrations(db)
	migrations.InitServices(db)
	migrations.InitTimeSlots(db)
	log.Print("migrations raised")
}
package main

import (
	"roombook_backend/internal/delivery"
	"roombook_backend/pkg/config"
	"roombook_backend/pkg/database"
	"roombook_backend/pkg/log"
)

func main() {
	log, loggerInfoFile, loggerErrorFile := log.InitLogger()

	defer loggerInfoFile.Close()
	defer loggerErrorFile.Close()

	config.InitConfig()
	log.Info("Config initialized")

	db := database.GetDB()
	log.Info("Database initialized")

	delivery.Start(db, log)

}

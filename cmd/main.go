package main

import (
	"backend_roombook/internal/delivery"
	"backend_roombook/pkg/config"
	"backend_roombook/pkg/database"
	"backend_roombook/pkg/log"
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

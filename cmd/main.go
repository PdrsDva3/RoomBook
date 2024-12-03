package main

import (
	"RoomBook/internal/delivery"
	"RoomBook/pkg/config"
	"RoomBook/pkg/database"
	"RoomBook/pkg/log"
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

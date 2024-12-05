package main

import (
	"RoomBook/internal/delivery"
	"RoomBook/pkg/config"
	"RoomBook/pkg/database"
	"RoomBook/pkg/database/cached"
	"RoomBook/pkg/log"
)

func main() {
	loger, loggerInfoFile, loggerErrorFile := log.InitLogger()

	defer loggerInfoFile.Close()
	defer loggerErrorFile.Close()

	config.InitConfig()
	loger.Info("Config initialized")

	db := database.GetDB()
	loger.Info("Database initialized")

	redis := cached.InitRedis()

	delivery.Start(db, loger, redis)

}

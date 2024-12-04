package main

import (
	"RoomBook/internal/delivery"
	"RoomBook/pkg/config"
	"RoomBook/pkg/database"
	"RoomBook/pkg/log"
)

const serviceName = "gin"

func main() {
	logger, loggerInfoFile, loggerErrorFile := log.InitLogger()

	defer loggerInfoFile.Close()
	defer loggerErrorFile.Close()

	config.InitConfig()
	logger.Info("Config initialized")

	//jaegerURL := fmt.Sprintf("http://%v:%v/api/traces", viper.GetString(config.JaegerHost), viper.GetString(config.JaegerPort))
	//tracer := trace.InitTracer(jaegerURL, serviceName)
	//logger.Info("Tracer Initialized")

	db := database.GetDB()
	logger.Info("Database initialized")

	//redisSession := cached.InitRedis(tracer)
	//logger.Info("Redis Initialized")

	delivery.Start(
		db,
		logger)
	//redisSession,
	//tracer)

}

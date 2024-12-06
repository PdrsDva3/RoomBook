package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

const (
	// данные для postgres
	DBName     = "DB_NAME"
	DBUser     = "DB_USER"
	DBPassword = "DB_PASSWORD"
	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"

	// данные для JWT
	TimeOut           = "TIME_OUT"
	JWTExpire         = "JWT_EXPIRE"
	Secret            = "SECRET"
	SessionExpiration = "SESSION_EXPIRATION"

	// данные для redis
	RedisHost     = "REDIS_HOST"
	RedisPassword = "REDIS_PASSWORD"
	RedisPort     = "REDIS_PORT"

	// данные для jaeger
	//JaegerHost = "JAEGER_HOST"
	//JaegerPort = "JAEGER_PORT"

	// данные для чата
	//DBMaxOpenConns = "DB_MAX_OPEN_CONNS"
)

func InitConfig() {
	envPath, _ := os.Getwd()
	envPath = filepath.Join(envPath, "..")
	envPath = filepath.Join(envPath, "/RoomBook/deploy")

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(envPath)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to init config. Error:%v", err.Error()))
	}
}

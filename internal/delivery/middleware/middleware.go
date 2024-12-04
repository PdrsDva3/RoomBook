package middleware

import (
	"RoomBook/pkg/log"
)

type Middleware struct {
	logger *log.Logs
	//jwtUtil auth.JWTUtil
	//session cached.Session
}

func InitMiddleware(logger *log.Logs) Middleware {
	return Middleware{
		logger: logger,
		//jwtUtil: util,
		//session: session,
	}
}

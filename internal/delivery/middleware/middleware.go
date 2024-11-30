package middleware

import (
	"RoomBook/pkg/auth"
	"RoomBook/pkg/database/cached"
	"RoomBook/pkg/log"
)

type Middleware struct {
	logger  *log.Logs
	jwtUtil auth.JWTUtil
	session cached.Session
}

func InitMiddleware(logger *log.Logs, util auth.JWTUtil, session cached.Session) Middleware {
	return Middleware{
		logger:  logger,
		jwtUtil: util,
		session: session,
	}
}

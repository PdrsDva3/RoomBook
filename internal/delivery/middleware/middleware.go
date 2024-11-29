package middleware

import (
	"roombook_backend/pkg/auth"
	"roombook_backend/pkg/database/cached"
	"roombook_backend/pkg/log"
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

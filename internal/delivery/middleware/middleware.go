package middleware

import (
	"roombook_backend/pkg/log"
)

type Middleware struct {
	logger *log.Logs
}

func InitMiddleware(logger *log.Logs) Middleware {
	return Middleware{
		logger: logger,
	}
}

package routers

import (
	"backend_roombook/internal/delivery/middleware"
	"backend_roombook/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRouting(r *gin.Engine, db *sqlx.DB, logger *log.Logs, middlewareStruct middleware.Middleware) {
	_ = RegisterUserRouter(r, db, logger)
}

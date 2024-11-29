package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"roombook_backend/internal/delivery/middleware"
	"roombook_backend/pkg/log"
)

func InitRouting(r *gin.Engine, db *sqlx.DB, logger *log.Logs, middlewareStruct middleware.Middleware) {
	_ = RegisterUserRouter(r, db, logger)
	_ = RegisterAdminRouter(r, db, logger)
	_ = RegisterHotelRouter(r, db, logger)
	_ = RegisterPhotoRouter(r, db, logger)
}

package routers

import (
	"RoomBook/internal/delivery/middleware"
	"RoomBook/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRouting(r *gin.Engine, db *sqlx.DB, logger *log.Logs, middlewareStruct middleware.Middleware) {
	_ = RegisterUserRouter(r, db, logger)
	_ = RegisterAdminRouter(r, db, logger)
	_ = RegisterHotelRouter(r, db, logger)
	_ = RegisterPhotoRouter(r, db, logger)
	_ = RegisterTagRouter(r, db, logger)
}

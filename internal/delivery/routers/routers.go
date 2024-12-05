package routers

import (
	"RoomBook/pkg/auth"
	"RoomBook/pkg/database/cached"
	"RoomBook/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRouting(r *gin.Engine, db *sqlx.DB, logger *log.Logs, jwt auth.JWTUtil, session cached.Session) {
	_ = RegisterUserRouter(r, db, logger, jwt, session) // middlewareStruct,
	_ = RegisterAdminRouter(r, db, logger)
	_ = RegisterHotelRouter(r, db, logger)
	_ = RegisterPhotoRouter(r, db, logger)
}

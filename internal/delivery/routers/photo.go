package routers

import (
	"RoomBook/internal/delivery/handlers"
	"RoomBook/internal/repository/photo"
	photoserv "RoomBook/internal/service/photo"
	"RoomBook/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterPhotoRouter(r *gin.Engine, db *sqlx.DB, logger *log.Logs) *gin.RouterGroup {
	photoRouter := r.Group("/photo")

	photoRepo := photo.InitPhotoRepository(db)
	photoService := photoserv.InitPhotoService(photoRepo, logger)
	photoHandler := handlers.InitPhotoHandler(photoService)

	photoRouter.POST("/add", photoHandler.Add)
	photoRouter.GET("/:id", photoHandler.Get)
	photoRouter.DELETE("/delete", photoHandler.Delete)

	return photoRouter
}

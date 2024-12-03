package routers

import (
	"RoomBook/internal/delivery/handlers"
	"RoomBook/internal/repository/hotel_photo"
	photoserv "RoomBook/internal/service/hotel_photo"
	"RoomBook/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterPhotoRouter(r *gin.Engine, db *sqlx.DB, logger *log.Logs) *gin.RouterGroup {
	photoRouter := r.Group("/hotel_photo")

	photoRepo := hotel_photo.InitPhotoRepository(db)
	photoService := photoserv.InitPhotoService(photoRepo, logger)
	photoHandler := handlers.InitPhotoHandler(photoService)

	photoRouter.POST("/add", photoHandler.Add)
	photoRouter.GET("/:id", photoHandler.Get)
	photoRouter.DELETE("/delete", photoHandler.Delete)

	return photoRouter
}

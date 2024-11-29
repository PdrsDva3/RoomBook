package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"roombook_backend/internal/delivery/handlers"
	"roombook_backend/internal/repository/hotel_photo"
	photoserv "roombook_backend/internal/service/hotel_photo"
	"roombook_backend/pkg/log"
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

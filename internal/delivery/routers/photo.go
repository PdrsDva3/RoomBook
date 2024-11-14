package routers

import (
	"backend_roombook/internal/delivery/handlers"
	"backend_roombook/internal/repository/photo"
	photoserv "backend_roombook/internal/service/photo"
	"backend_roombook/pkg/log"
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

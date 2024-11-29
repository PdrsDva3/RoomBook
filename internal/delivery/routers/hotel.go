package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"roombook_backend/internal/delivery/handlers"
	"roombook_backend/internal/repository/hotel"
	hotelserv "roombook_backend/internal/service/hotel"
	"roombook_backend/pkg/log"
)

func RegisterHotelRouter(r *gin.Engine, db *sqlx.DB, logger *log.Logs) *gin.RouterGroup {
	hotelRouter := r.Group("/hotel")

	hotelRepo := hotel.InitHotelRepository(db)
	hotelService := hotelserv.InitHotelService(hotelRepo, logger)
	hotelHandler := handlers.InitHotelHandlers(hotelService)

	hotelRouter.POST("/create", hotelHandler.Create)
	hotelRouter.GET("/:id", hotelHandler.Get)
	hotelRouter.DELETE("/delete/:id", hotelHandler.Delete)

	return hotelRouter
}

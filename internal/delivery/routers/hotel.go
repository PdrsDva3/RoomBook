package routers

import (
	"backend_roombook/internal/delivery/handlers"
	"backend_roombook/internal/repository/hotel"
	hotelserv "backend_roombook/internal/service/hotel"
	"backend_roombook/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterHotelRouter(r *gin.Engine, db *sqlx.DB, logger *log.Logs) *gin.RouterGroup {
	hotelRouter := r.Group("/hotel")

	hotelRepo := hotel.InitHotelRepository(db)
	hotelService := hotelserv.InitHotelService(hotelRepo, logger)
	hotelHandler := handlers.InitHotelHandlers(hotelService)

	hotelRouter.POST("/create", hotelHandler.Create)
	hotelRouter.GET("/:id", hotelHandler.Get)
	hotelRouter.DELETE("/:id", hotelHandler.Delete)

	return hotelRouter
}

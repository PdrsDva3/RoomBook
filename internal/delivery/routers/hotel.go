package routers

import (
	"RoomBook/internal/delivery/handlers"
	"RoomBook/internal/repository/hotel"
	hotelserv "RoomBook/internal/service/hotel"
	"RoomBook/pkg/log"
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
	hotelRouter.DELETE("/delete/:id", hotelHandler.Delete)

	return hotelRouter
}

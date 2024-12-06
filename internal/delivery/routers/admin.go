package routers

import (
	"RoomBook/internal/delivery/handlers"
	"RoomBook/internal/repository/admin"
	adminserv "RoomBook/internal/service/admin"
	"RoomBook/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterAdminRouter(r *gin.Engine, db *sqlx.DB, logger *log.Logs) *gin.RouterGroup {
	adminRouter := r.Group("/admin")

	adminRepo := admin.InitAdminRepository(db)
	adminService := adminserv.InitAdminService(adminRepo, logger)
	adminHandler := handlers.InitAdminHandler(adminService)

	adminRouter.POST("/create", adminHandler.Create)
	adminRouter.POST("/login", adminHandler.Login)
	adminRouter.GET("/:id", adminHandler.Get)
	adminRouter.DELETE("/delete/:id", adminHandler.Delete)

	return adminRouter
}

package routers

import (
	"backend_roombook/internal/delivery/handlers"
	"backend_roombook/internal/repository/admin"
	adminserv "backend_roombook/internal/service/admin"
	"backend_roombook/pkg/log"
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
	adminRouter.PUT("/change/pwd", adminHandler.ChangePWD)
	adminRouter.DELETE("/delete/:id", adminHandler.Delete)

	return adminRouter
}

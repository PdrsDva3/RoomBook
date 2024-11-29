package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"roombook_backend/internal/delivery/handlers"
	"roombook_backend/internal/repository/admin"
	adminserv "roombook_backend/internal/service/admin"
	"roombook_backend/pkg/log"
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

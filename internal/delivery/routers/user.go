package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"roombook_backend/internal/delivery/handlers"
	"roombook_backend/internal/repository/user"
	userserv "roombook_backend/internal/service/user"
	"roombook_backend/pkg/log"
)

func RegisterUserRouter(r *gin.Engine, db *sqlx.DB, logger *log.Logs) *gin.RouterGroup {
	userRouter := r.Group("/user")

	userRepo := user.InitUserRepository(db)
	userService := userserv.InitUserService(userRepo, logger)
	userHandler := handlers.InitUserHandler(userService)

	userRouter.POST("/create", userHandler.Create)
	userRouter.POST("/login", userHandler.Login)
	userRouter.GET("/:id", userHandler.Get)
	userRouter.PUT("/change/pwd", userHandler.ChangePWD)
	userRouter.DELETE("/delete/:id", userHandler.Delete)
	return userRouter
}

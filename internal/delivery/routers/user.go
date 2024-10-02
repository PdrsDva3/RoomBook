package routers

import (
	"backend_roombook/internal/delivery/handlers"
	"backend_roombook/internal/repository/user"
	userserv "backend_roombook/internal/service/user"
	"backend_roombook/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
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

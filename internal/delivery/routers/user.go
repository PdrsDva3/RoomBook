package routers

import (
	"RoomBook/internal/delivery/handlers"
	"RoomBook/internal/repository/user"
	userserv "RoomBook/internal/service/user"
	"RoomBook/pkg/auth"
	"RoomBook/pkg/database/cached"
	"RoomBook/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterUserRouter(r *gin.Engine, db *sqlx.DB, logger *log.Logs, jwt auth.JWTUtil, session cached.Session) *gin.RouterGroup {
	userRouter := r.Group("/user")

	userRepo := user.InitUserRepository(db)
	userService := userserv.InitUserService(userRepo, logger, jwt, session)
	userHandler := handlers.InitUserHandler(userService)

	userRouter.POST("/create", userHandler.Create)
	userRouter.POST("/login", userHandler.Login)
	userRouter.GET("/:id", userHandler.Get)
	//userRouter.PUT("/change/pwd", userHandler.ChangePWD)
	userRouter.DELETE("/delete/:id", userHandler.Delete)
	return userRouter
}

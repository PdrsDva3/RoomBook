package delivery

import (
	"RoomBook/cmd/docs"
	"RoomBook/internal/delivery/middleware"
	"RoomBook/internal/delivery/routers"
	"RoomBook/pkg/log"
	"RoomBook/pkg/auth"
	"RoomBook/pkg/database/cached"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	//"go.opentelemetry.io/otel/trace"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start(db *sqlx.DB, log *log.Logs, session cached.Session) {
	r := gin.Default()
	r.ForwardedByClientIP = true

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	jwtUtils := auth.InitJWTUtil()
	middlewareStruct := middleware.InitMiddleware(log, jwtUtils, session)
	r.Use(middlewareStruct.CORSMiddleware())

	routers.InitRouting(r, db, log, jwtUtils, session) // jwtUtils, session, tracer

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}

package delivery

import (
	"RoomBook/docs"
	"RoomBook/internal/delivery/middleware"
	"RoomBook/internal/delivery/routers"
	"RoomBook/pkg/auth"
	"RoomBook/pkg/database/cached"
	"RoomBook/pkg/log"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel/trace"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start(db *sqlx.DB, log *log.Logs, session cached.Session, tracer trace.Tracer) {
	r := gin.Default()
	r.ForwardedByClientIP = true

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	jwtUtils := auth.InitJWTUtil()
	middlewareStruct := middleware.InitMiddleware(log, jwtUtils, session)
	r.Use(middlewareStruct.CORSMiddleware())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"} // Добавьте нужные вам домены
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	corsConfig.MaxAge = 12 * time.Hour

	router.Use(cors.New(corsConfig))

	routers.InitRouting(r, db, log, middlewareStruct) // jwtUtils, session, tracer
	tt := 12 * time.Hour
	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}

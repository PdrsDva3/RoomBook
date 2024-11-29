package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel/trace"
	"roombook_backend/docs"
	"roombook_backend/internal/delivery/middleware"
	"roombook_backend/internal/delivery/routers"
	"roombook_backend/pkg/auth"
	"roombook_backend/pkg/database/cached"
	"roombook_backend/pkg/log"

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

	routers.InitRouting(r, db, log, middlewareStruct) // jwtUtils, session, tracer

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}

package routers

import (
	"RoomBook/internal/delivery/handlers"
	"RoomBook/internal/repository/tag"
	tagserv "RoomBook/internal/service/tag"
	"RoomBook/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterTagRouter(r *gin.Engine, db *sqlx.DB, logger *log.Logs) *gin.RouterGroup {
	tagRouter := r.Group("/tag")

	tagRepo := tag.InitTagRepository(db)
	tagService := tagserv.InitTagService(tagRepo, logger)
	tagHandler := handlers.InitTagHandler(tagService)

	tagRouter.POST("/hotel", tagHandler.AddHotel)
	tagRouter.POST("/room", tagHandler.AddRoom)
	tagRouter.DELETE("/hotel/:id_hotel/:id_tag", tagHandler.DeleteHotel)
	tagRouter.DELETE("/room/:id_room/:id_tag", tagHandler.DeleteRoom)

	tagTypeRepo := tag.InitTagTypeRepository(db)
	tagTypeServ := tagserv.InitTagTypeService(tagTypeRepo, logger)
	tagTypeHandler := handlers.InitTagTypeHandler(tagTypeServ)

	tagRouter.POST("/create", tagTypeHandler.CreateTag)
	tagRouter.POST("/type/create", tagTypeHandler.CreateType)
	tagRouter.GET("/", tagTypeHandler.GetTags)
	tagRouter.GET("/type", tagTypeHandler.GetType)
	tagRouter.GET("/type/:id_type", tagTypeHandler.GetTagType)

	return tagRouter
}

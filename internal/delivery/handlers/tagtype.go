package handlers

import (
	"RoomBook/internal/models"
	"RoomBook/internal/service"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TagTypeHandler struct {
	service service.TagTypeServ
}

func InitTagTypeHandler(service service.TagTypeServ) TagTypeHandler {
	return TagTypeHandler{
		service: service,
	}
}

// @Summary create tag
// @Tags tag
// @Accept  json
// @Produce  json
// @Param data body models.TagCreate true "tag"
// @Success 200 {object} int "Successfully created tag"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tag/create [post]
func (h TagTypeHandler) CreateTag(g *gin.Context) {
	var newTag models.TagCreate

	if err := g.ShouldBindJSON(&newTag); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := h.service.CreateTag(ctx, newTag)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary create type
// @Tags tag
// @Accept  json
// @Produce  json
// @Param data body models.TypeCreate true "type"
// @Success 200 {object} int "Successfully created tag"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tag/type/create [post]
func (h TagTypeHandler) CreateType(g *gin.Context) {
	var newType models.TypeCreate

	if err := g.ShouldBindJSON(&newType); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := h.service.CreateType(ctx, newType)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Get tag
// @Tags tag
// @Accept  json
// @Produce  json
// @Success 200 {object} int "Successfully get"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tag/ [get]
func (h TagTypeHandler) GetTags(g *gin.Context) {
	ctx := g.Request.Context()

	tags, err := h.service.Tags(ctx)
	if err != nil {
		g.JSON(http.StatusTeapot, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"tags": tags})
}

// @Summary Get type
// @Tags tag
// @Accept  json
// @Produce  json
// @Success 200 {object} int "Successfully get admin"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tag/type [get]
func (h TagTypeHandler) GetType(g *gin.Context) {
	ctx := g.Request.Context()

	types, err := h.service.Types(ctx)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"types": types})
}

// @Summary Get tag type
// @Tags tag
// @Accept  json
// @Produce  json
// @Param id_type query int true "TypeID"
// @Success 200 {object} int "Successfully get admin"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tag/type/{id_type} [get]
func (h TagTypeHandler) GetTagType(g *gin.Context) {
	id := g.Query("id_type")
	aid, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx := g.Request.Context()

	tags, err := h.service.TagsType(ctx, aid)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"tags": tags})
}

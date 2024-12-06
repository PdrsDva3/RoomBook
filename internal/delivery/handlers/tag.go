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

type TagHandler struct {
	service service.TagServ
}

func InitTagHandler(service service.TagServ) TagHandler {
	return TagHandler{
		service: service,
	}
}

// @Summary Add hotel tag
// @Tags tag
// @Accept  json
// @Produce  json
// @Param data body models.TagHotel true "tag hotel"
// @Success 200 {object} int "Successfully created tag"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tag/hotel [post]
func (h TagHandler) AddHotel(g *gin.Context) {
	var newTag models.TagHotel

	if err := g.ShouldBindJSON(&newTag); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := h.service.AddHotel(ctx, newTag)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"add": "success"})
}

// @Summary Add room tag
// @Tags tag
// @Accept  json
// @Produce  json
// @Param data body models.TagRoom true "tag room"
// @Success 200 {object} int "Successfully created tag"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tag/room [post]
func (h TagHandler) AddRoom(g *gin.Context) {
	var newTag models.TagRoom

	if err := g.ShouldBindJSON(&newTag); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := h.service.AddRoom(ctx, newTag)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"add": "success"})
}

// @Summary Delete hotel tag
// @Tags tag
// @Accept  json
// @Produce  json
// @Param id_hotel query int true "HotelID"
// @Param id_tag query int true "TagID"
// @Success 200 {object} int "Successfully created tag"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tag/hotel/{id_hotel}/{id_tag} [delete]
func (h TagHandler) DeleteHotel(g *gin.Context) {
	var tag models.TagHotel
	id := g.Query("id_hotel")

	aid, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tag.IDHotel = aid

	id = g.Query("id_tag")

	aid, err = strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tag.IDTag = aid

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = h.service.DeleteHotel(ctx, tag)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"delete": "success"})
}

// @Summary Delete room tag
// @Tags tag
// @Accept  json
// @Produce  json
// @Param id_room query int true "RoomID"
// @Param id_tag query int true "TagID"
// @Success 200 {object} int "Successfully created tag"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tag/room/{id_room}/{id_tag} [delete]
func (h TagHandler) DeleteRoom(g *gin.Context) {
	var tag models.TagRoom
	id := g.Query("id_room")

	aid, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tag.IDRoom = aid

	id = g.Query("id_tag")

	aid, err = strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tag.IDTag = aid

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = h.service.DeleteRoom(ctx, tag)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"delete": "success"})
}

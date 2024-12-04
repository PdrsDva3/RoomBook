package handlers

import (
	"RoomBook/internal/models"
	"RoomBook/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type PhotoHandler struct {
	service service.PhotoServ
}

func InitPhotoHandler(service service.PhotoServ) PhotoHandler {
	return PhotoHandler{service: service}
}

// @Summary Add hotel_photo
// @Tags hotel_photo
// @Accept  json
// @Produce  json
// @Param data body models.PhotoAddWithIDHotel true "add Photo"
// @Success 200 {object} int "Successfully add"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /hotel/photo/ [post]
func (h PhotoHandler) Add(g *gin.Context) {
	var newPhotos models.PhotoAddWithIDHotel

	if err := g.ShouldBindJSON(&newPhotos); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := h.service.Add(ctx, newPhotos)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"out": "success"})
}

// @Summary Get hotel_photo
// @Tags hotel_photo
// @Accept  json
// @Produce  json
// @Param id query int true "HotelID"
// @Success 200 {object} int "Successfully get"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /hotel/photo/{id} [get]
func (h PhotoHandler) Get(g *gin.Context) {
	id := g.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx := g.Request.Context()

	photo, err := h.service.Get(ctx, aid)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"hotel_photo": photo})
}

// @Summary Delete photos
// @Tags hotel_photo
// @Accept  json
// @Produce  json
// @Param data body models.PhotoDelete true "PhotoIDs"
// @Success 200 {object} int "Successfully deleted"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /hotel/photo/ [delete]
func (h PhotoHandler) Delete(g *gin.Context) {
	var photo models.PhotoDelete
	if err := g.ShouldBindJSON(&photo); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := h.service.Delete(ctx, photo.ID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"delete": "success"})
}

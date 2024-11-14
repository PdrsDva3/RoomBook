package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"roombook_backend/internal/models"
	"roombook_backend/internal/service"
	"strconv"
	"time"
)

type HotelHandlers struct {
	service service.HotelServ
}

func InitHotelHandlers(service service.HotelServ) HotelHandlers {
	return HotelHandlers{service: service}
}

// @Summary Create hotel
// @Tags hotel
// @Accept  json
// @Produce  json
// @Param data body models.HotelCreate true "hotel create"
// @Success 200 {object} int "Successfully created"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /hotel/create [post]
func (h HotelHandlers) Create(g *gin.Context) {
	var newHotel models.HotelCreate
	if err := g.ShouldBindJSON(&newHotel); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := h.service.Create(ctx, newHotel)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Get hotel
// @Tags hotel
// @Accept  json
// @Produce  json
// @Param id query int true "HotelID"
// @Success 200 {object} int "Successfully get"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /hotel/{id} [get]
func (h HotelHandlers) Get(g *gin.Context) {
	id := g.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := g.Request.Context()

	hotel, err := h.service.Get(ctx, aid)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"hotel": hotel})
}

// @Summary Delete hotel
// @Tags hotel
// @Accept  json
// @Produce  json
// @Param id query int true "HotelID"
// @Success 200 {object} int "Successfully deleted"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /hotel/delete/{id} [delete]
func (h HotelHandlers) Delete(g *gin.Context) {
	id := g.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := g.Request.Context()
	err = h.service.Delete(ctx, aid)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"hotel": nil})
}

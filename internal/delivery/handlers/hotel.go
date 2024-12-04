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
// @Router /hotel/ [post]
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

// @Summary GetAll hotel
// @Tags hotel
// @Accept  json
// @Produce  json
// @Success 200 {object} int "Successfully get"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /hotel/ [get]
func (h HotelHandlers) GetAll(g *gin.Context) {
	ctx := g.Request.Context()

	hotels, err := h.service.GetAll(ctx)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"hotels": hotels})
}

// @Summary Change hotel
// @Tags hotel
// @Accept  json
// @Produce  json
// @Param data body models.HotelChange true "change hotel"
// @Success 200 {object} int "Success changing"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /hotel/ [put]
func (h HotelHandlers) Change(g *gin.Context) {
	var hotel models.HotelChange
	if err := g.ShouldBindJSON(&hotel); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := h.service.Change(ctx, hotel)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": hotel.IDHotel})
}

// @Summary Add hotel admin
// @Tags hotel
// @Accept  json
// @Produce  json
// @Param idHotel query int true "HotelID"
// @Param idAdmin query int true "AdminID"
// @Success 200 {object} int "Successfully created"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /hotel/{idHotel}/{idAdmin} [post]
func (h HotelHandlers) Admin(g *gin.Context) {
	hotelID := g.Query("idHotel")
	adminID := g.Query("idAdmin")
	id1, err := strconv.Atoi(hotelID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id2, err := strconv.Atoi(adminID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	err = h.service.Admin(ctx, models.HotelAdmin{id1, id2})
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"add": "success"})
}

// @Summary Rating hotel
// @Tags hotel
// @Accept  json
// @Produce  json
// @Param data body models.HotelRating true "hotel rating"
// @Success 200 {object} int "Successfully created"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /hotel/rating [post]
func (h HotelHandlers) Rating(g *gin.Context) {
	var Hotel models.HotelRating
	if err := g.ShouldBindJSON(&Hotel); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := h.service.Rating(ctx, Hotel)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"rating": "success"})
}

// @Summary Delete hotel
// @Tags hotel
// @Accept  json
// @Produce  json
// @Param id query int true "HotelID"
// @Success 200 {object} int "Successfully deleted"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /hotel/{id} [delete]
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

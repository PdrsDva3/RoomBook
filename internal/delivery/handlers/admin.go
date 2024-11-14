package handlers

import (
	"backend_roombook/internal/models"
	"backend_roombook/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type AdminHandler struct {
	service service.AdminServ
}

func InitAdminHandler(service service.AdminServ) AdminHandler {
	return AdminHandler{service: service}
}

// @Summary Create admin
// @Tags admin
// @Accept  json
// @Produce  json
// @Param data body models.AdminCreate true "admin create"
// @Success 200 {object} int "Successfully created admin"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/create [post]
func (h AdminHandler) Create(g *gin.Context) {
	var newAdmin models.AdminCreate

	if err := g.ShouldBindJSON(&newAdmin); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := h.service.Create(ctx, newAdmin)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Get admin
// @Tags admin
// @Accept  json
// @Produce  json
// @Param id query int true "AdminID"
// @Success 200 {object} int "Successfully get admin"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/{id} [get]
func (h AdminHandler) Get(g *gin.Context) {
	id := g.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx := g.Request.Context()

	admin, err := h.service.Get(ctx, aid)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"admin": admin})
}

// @Summary Change password
// @Tags admin
// @Accept  json
// @Produce  json
// @Param data body models.AdminChangePWD true "change password"
// @Success 200 {object} int "Success changing"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/change/pwd [put]
func (h AdminHandler) ChangePWD(g *gin.Context) {
	var admin models.AdminChangePWD
	if err := g.ShouldBindJSON(&admin); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := h.service.ChangePWD(ctx, admin)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Login admin
// @Tags admin
// @Accept  json
// @Produce  json
// @Param data body models.AdminLogin true "admin login"
// @Success 200 {object} int "Successfully login admin"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/login [post]
func (h AdminHandler) Login(g *gin.Context) {
	var admin models.AdminLogin
	if err := g.ShouldBindJSON(&admin); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := g.Request.Context()

	id, err := h.service.Login(ctx, admin)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Delete admin
// @Tags admin
// @Accept  json
// @Produce  json
// @Param id query int true "AdminID"
// @Success 200 {object} int "Successfully deleted"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/delete/{id} [delete]
func (h AdminHandler) Delete(g *gin.Context) {
	id := g.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = h.service.Delete(ctx, aid)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

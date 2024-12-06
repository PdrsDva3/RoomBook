package handlers

import (
	"RoomBook/internal/models"
	"RoomBook/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserChangeHandler struct {
	service service.UserChangeServ
}

func InitUserChangeHandler(service service.UserChangeServ) UserChangeHandler {
	return UserChangeHandler{
		service: service,
	}
}

// @Summary Change password
// @Tags user
// @Accept  json
// @Produce  json
// @Param data body models.UserChangePWD true "change password"
// @Success 200 {object} int "Success changing"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/change/pwd [put]
func (h UserChangeHandler) ChangePWD(g *gin.Context) {
	var user models.UserChangePWD
	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := h.service.PWD(ctx, user)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"change": "success", "id": ""})
}

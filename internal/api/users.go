package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"reit-real-estate/internal/dto"
)

type userService interface {
	RegisterUser(ctx context.Context, dto *dto.RegisterUserDTO) error
}

func (c *controller) SignUp(ctx *gin.Context) {
	var request *dto.RegisterUserDTO
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.userService.RegisterUser(ctx, request)
	if err != nil {
		//TODO handle errors
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

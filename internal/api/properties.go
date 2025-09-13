package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reit-real-estate/internal/dto"
)

func (c *controller) RegisterProperty(ctx *gin.Context) {
	var request *dto.RegisterPropertyDTO
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.service.RegisterProperty(ctx, request)
	if err != nil {
		//TODO handle errors
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (c *controller) Invest(ctx *gin.Context) {
	var request *dto.InvestDTO
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.service.Invest(ctx, request)
	if err != nil {
		//TODO handle errors
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

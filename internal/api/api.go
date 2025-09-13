package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"reit-real-estate/internal/dto"
)

type service interface {
	RegisterUser(ctx context.Context, dto *dto.RegisterUserDTO) error
	RegisterProperty(ctx context.Context, request *dto.RegisterPropertyDTO) error
}

type controller struct {
	service service
}

func NewController(service service) *controller {
	return &controller{
		service: service,
	}
}

func (c *controller) Routes(r *gin.Engine) {
	r.POST("/v1/users", c.SignUp)
	r.POST("/v1/properties", c.RegisterProperty)
}

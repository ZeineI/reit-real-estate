package api

import (
	"github.com/gin-gonic/gin"
)

type controller struct {
	userService userService
}

func NewController(userService userService) *controller {
	return &controller{
		userService: userService,
	}
}

func (c *controller) Routes(r *gin.Engine) {
	r.POST("/v1/users", c.SignUp)
}

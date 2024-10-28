package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (uc *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", " fullstack dev")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...pingdangth" + name,
		"uid":     uid,
		"user":    []string{"newman", "jackson", "weston&smith"},
	})
}

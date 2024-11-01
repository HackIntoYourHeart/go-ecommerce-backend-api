package routers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	var r *gin.Engine
	return r
}

// func PongWithParams(params string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		name := c.DefaultQuery("name", "fullstack dev")
// 		uid := c.Query("uid")

// 		// Truy cập tham số params từ closure
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong...pingdangth " + name + " " + params,
// 			"uid":     uid,
// 		})
// 	}
// }

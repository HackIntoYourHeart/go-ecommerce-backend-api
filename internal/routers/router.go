package routers

import (
	"net/http"

	c "example.com/go-ecommerce-backend-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	// var result string = "test"

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user/1", c.NewUserController().GetUserById)
		// v1.PUT("/ping", Pong)
		// v1.PATCH("/ping", Pong)
		// v1.DELETE("/ping", Pong)
		// v1.HEAD("/ping", Pong)
		// v1.OPTIONS("/ping", Pong)
	}

	// v2 := r.Group("/v2/2024")
	// {
	// 	v2.GET("/ping", Pong)
	// 	v2.PUT("/ping", Pong)
	// 	v2.PATCH("/ping", Pong)
	// 	v2.DELETE("/ping", Pong)
	// 	v2.HEAD("/ping", Pong)
	// 	v2.OPTIONS("/ping", Pong)
	// }
	return r
}

func PongWithParams(params string) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.DefaultQuery("name", "fullstack dev")
		uid := c.Query("uid")

		// Truy cập tham số params từ closure
		c.JSON(http.StatusOK, gin.H{
			"message": "pong...pingdangth " + name + " " + params,
			"uid":     uid,
		})
	}
}

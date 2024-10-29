package middlewares

import (
	"strings"

	"example.com/go-ecommerce-backend-api/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c * gin.Context){
		token := c.GetHeader("Authorization")
		if strings.TrimPrefix(token, "Bearer ") != "valid-token" {
			response.ErrorResponse(c, response.ErrCodeInvalidToken, "Invalid token")
			c.Abort()
			return
		}
	}
}
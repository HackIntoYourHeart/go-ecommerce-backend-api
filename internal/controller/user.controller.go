package controller

import (
	"example.com/go-ecommerce-backend-api/internal/service"
	"example.com/go-ecommerce-backend-api/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	// name := c.DefaultQuery("name", " fullstack dev")
	// uid := c.Query("uid")
	// if err != nil {
	// 	return response.ErrorResponse(c, 20003, "Something wrong!")
	// }
	// response.ErrorResponse(c, 20003, "Something wrong!")
	response.SuccessResponse(c, 20001, []string{"nodejs", "golang", "java", "reactjs"})
	
}

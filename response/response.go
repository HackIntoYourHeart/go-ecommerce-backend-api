package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`    // status code
	Message string      `json:"message"` // error message
	Data    interface{} `json:"data"`   // data return
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message interface{}) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
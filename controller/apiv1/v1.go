package apiv1

import (
	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, code HttpStatus, data interface{}) {
	c.JSON(code.Code, gin.H{
		"code": code.Code,
		"msg":  code.Msg,
		"data": data,
	})
}

var ExampleErrorResponse = ErrorResponse{500, "Internal error"}

type ErrorResponse struct {
	Code int
	Msg  string
}

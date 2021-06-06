package apiv1

import (
	"singlishwords/controller/api/v1/status"

	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, code status.Code, data interface{}) {
	c.JSON(code.Code, gin.H{
		"code": code.Code,
		"msg":  code.Msg,
		"data": data,
	})
}

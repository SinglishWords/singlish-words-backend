package apiv1

import (
	"github.com/gin-gonic/gin"
	"singlishwords/config"
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

func ShowAllSub(c *gin.Context) (HttpStatus, interface{}) {
	return StatusOK, gin.H{
		"answers_url":     config.App.BaseURL + "answers" + "{?limit}",
		"questions_url":   config.App.BaseURL + "questions" + "{?limit}",
		"respondents_url": config.App.BaseURL + "respondents" + "{?limit}",
	}
}

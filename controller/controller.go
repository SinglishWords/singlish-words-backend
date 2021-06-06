package controller

import (
	apiv1 "singlishwords/controller/api/v1"
	"singlishwords/controller/api/v1/status"

	"github.com/gin-gonic/gin"
)

// type handler func(*gin.Context) (status.Code, error)

func ResponseWrapper(f func(*gin.Context) (status.Code, interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, data := f(c)
		response := gin.H{
			"code": code.Code,
			"msg":  code.Msg,
		}
		if data != nil {
			response["data"] = data
		}
		c.JSON(code.Code, response)
	}
}

func InitRouter(g *gin.Engine) error {
	v1 := g.Group("/api/v1")
	{
		// Questions
		v1.GET("/questions", ResponseWrapper(apiv1.GetQuestions))

		// Answers
		v1.GET("/answers", ResponseWrapper(apiv1.GetAnswers))
		v1.POST("/answers", ResponseWrapper(apiv1.PostAnswers))
	}
	return nil
}

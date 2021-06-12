package router

import (
	"github.com/gin-gonic/gin"
	"singlishwords/controller/apiv1"
	"singlishwords/controller/apiv1/answer"
	"singlishwords/controller/apiv1/question"
	"singlishwords/controller/apiv1/respondent"
)

func InitRouter(g *gin.Engine) *gin.Engine {

	v1 := g.Group("/api/v1")
	{
		// Question
		v1.GET("/question", responseWrapper(question.GetQuestions))

		// Answer
		v1.GET("/answer", responseWrapper(answer.GetAnswers))
		v1.POST("/answer", responseWrapper(answer.PostAnswer))
		// v1.POST("/answer", ResponseWrapper(answerAPI.PostAnswer))

		// Respondent
		v1.POST("/respondent", responseWrapper(respondent.PostRespondent))
		v1.GET("/respondent", responseWrapper(respondent.GetRespondent))
	}

	return g
}

func responseWrapper(f func(*gin.Context) (apiv1.HttpStatus, interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, data := f(c)
		if code != apiv1.StatusOK {
			c.JSON(code.Code, gin.H{
				"code":    code.Code,
				"message": code.Msg,
			})
		} else {
			c.JSON(code.Code, data)
		}
	}
}

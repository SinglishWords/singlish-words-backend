package router

import (
	"github.com/gin-gonic/gin"
	"singlishwords/config"
	"singlishwords/controller/apiv1"
	"singlishwords/controller/apiv1/answer"
	"singlishwords/controller/apiv1/question"
	"singlishwords/controller/apiv1/respondent"
	"singlishwords/log"
)

func InitRouter(g *gin.Engine) *gin.Engine {
	g.Use(log.RouteLogger())
	v1 := g.Group(config.App.BaseURL)
	{
		v1.GET("/", responseWrapper(apiv1.ShowAllSub))
		// Question
		v1.GET("/questions", responseWrapper(question.GetQuestions))

		// Answer
		v1.GET("/answers", responseWrapper(answer.GetAnswers))
		v1.POST("/answer", responseWrapper(answer.PostAnswer))
		// v1.POST("/answer", ResponseWrapper(answerAPI.PostAnswer))

		// Together
		v1.POST("/answers", responseWrapper(answer.PostRespondentWithAnswers))

		// Respondent
		v1.POST("/respondent", responseWrapper(respondent.PostRespondent))
		v1.GET("/respondents", responseWrapper(respondent.GetRespondent))
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

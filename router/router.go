package router

import (
	"github.com/gin-gonic/gin"
	"singlishwords/config"
	"singlishwords/controller/apiv1"
	"singlishwords/controller/apiv1/answer"
	"singlishwords/controller/apiv1/email"
	"singlishwords/controller/apiv1/question"
	"singlishwords/controller/apiv1/respondent"
	"singlishwords/middleware"
)

func InitRouter(g *gin.Engine) *gin.Engine {
	g.Use(middleware.RouteLogger())
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
		v1.GET("/respondents", responseWrapper(respondent.GetRespondent))
		v1.POST("/respondent", responseWrapper(respondent.PostRespondent))
		//v1.PATCH("/respondent", responseWrapper(respondent.PatchRespondent))

		// Email
		v1.POST("/email", responseWrapper(email.PostEmail))
	}

	return g
}

func responseWrapper(f func(*gin.Context) (apiv1.HttpStatus, interface{})) gin.HandlerFunc {
	returnWithContentStatus := map[apiv1.HttpStatus]struct{}{
		apiv1.StatusOK:      {},
		apiv1.StatusCreated: {},
	}

	return func(c *gin.Context) {
		code, data := f(c)
		if _, ok := returnWithContentStatus[code]; ok {
			c.JSON(code.Code, data)
		} else {
			c.JSON(code.Code, gin.H{
				"code":    code.Code,
				"message": code.Msg,
			})
		}
	}
}

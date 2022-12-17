package router

import (
	"singlishwords/config"
	"singlishwords/controller/apiv1"
	"singlishwords/controller/apiv1/answer"
	"singlishwords/controller/apiv1/association"
	"singlishwords/controller/apiv1/email"
	"singlishwords/controller/apiv1/question"
	"singlishwords/controller/apiv1/respondent"
	"singlishwords/middleware"

	"github.com/gin-gonic/gin"
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

		// Association
		v1.GET("/associations/forward/:word", responseWrapper(association.GetForwardAssociations))
		v1.GET("/associations/backward/:word", responseWrapper(association.GetBackwardAssociations))
		// v1.GET("/associations/random/forward", responseWrapper(association.GetForwardAssociations))
		// v1.GET("/associations/random/backward", responseWrapper(association.GetForwardAssociations))
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

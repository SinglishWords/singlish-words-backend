package question

import (
	"singlishwords/controller/apiv1"
	"singlishwords/service"

	"strconv"

	"github.com/gin-gonic/gin"
)

// param name,param type,data type,is mandatory?,comment attribute(optional)
// return code,{param type},data type,comment

// GetQuestions godoc
// @Summary Get a list of questions
// @Description Get a list of questions
// @Tags Questions
// @Produce json
// @Param limit query int false "Will return limit number of questions, default=8"
// @Success 200 {object} []model.Question
// @Failure 500 {object} apiv1.ErrorResponse
// @Router /api/v1/question [get]
func GetQuestions(c *gin.Context) (apiv1.HttpStatus, interface{}) {
	code := apiv1.StatusOK

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "8"))
	if err != nil {
		code = apiv1.StatusQueryParamError
		return code, nil
	}

	questions, err := service.GetRandomNQuestions(limit)
	if err != nil {
		// code = status.Fail("Failed when retrieve data from database.")
		code = apiv1.StatusFail(err.Error())
		return code, nil
	}

	return code, questions
}

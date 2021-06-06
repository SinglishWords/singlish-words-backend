package apiv1

import (
	"singlishwords/controller/api/v1/status"
	"singlishwords/model"
	question_service "singlishwords/service/question"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetQuestions godoc
// @Summary Get a list of questions
// @Description Get a list of questions
// @Tags Questions
// @Produce json
// @Param limit query string false "Will return limit number of questions." Enums(A, B, C)
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /api/v1/questions [get]
func GetQuestions(c *gin.Context) (code status.Code, questions interface{}) {
	code, questions, err := status.Success, []model.Question{}, error(nil)

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "8"))
	if err != nil {
		code = status.Fail("Query parameter limit should be int.")
		return
	}

	questions, err = question_service.GetQuestions(limit)
	if err != nil {
		// code = status.Fail("Failed when retrive data from database.")
		code = status.Fail(err.Error())
		return
	}

	return
}

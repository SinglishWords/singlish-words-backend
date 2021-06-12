package answer

import (
	"github.com/gin-gonic/gin"
	"singlishwords/controller/apiv1"
	"singlishwords/service"
)

// GetAnswers godoc
// @Summary Get all answers
// @Tags Answer
// @Produce json
// @Param limit query int false "default=10000"
// @Success 201 {object} []model.Answer
// @Failure 500 {object} apiv1.ErrorResponse
// @Router /answer [get]
func GetAnswers(c *gin.Context) (apiv1.HttpStatus, interface{}) {
	param := struct {
		Limit int
	}{
		Limit: 10000,
	}
	err := c.BindQuery(&param)
	if err != nil {
		return apiv1.StatusQueryParamError, nil
	}
	answers, err := service.GetAnswers(param.Limit)
	if err != nil {
		return apiv1.StatusFail(err.Error()), err
	}
	return apiv1.StatusOK, answers
}

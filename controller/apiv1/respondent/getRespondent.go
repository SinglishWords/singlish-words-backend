package respondent

import (
	"github.com/gin-gonic/gin"
	"singlishwords/controller/apiv1"
	"singlishwords/service"
)

// GetRespondent godoc
// @Summary Get all respondents
// @Tags Respondent
// @Produce json
// @Param limit query int false "default=10000"
// @Success 200 {object} []model.Respondent
// @Failure 500 {object} apiv1.ErrorResponse
// @Router /respondent [get]
func GetRespondent(c *gin.Context) (apiv1.HttpStatus, interface{}) {
	code := apiv1.StatusOK
	param := struct {
		Limit int `json:"limit"`
	}{
		Limit: 10000,
	}

	err := c.BindQuery(&param)
	if err != nil {
		code = apiv1.StatusQueryParamError
		return code, nil
	}

	respondents, err := service.GetAllRespondents()

	if err != nil {
		code = apiv1.StatusFail(err.Error())
		return code, nil
	}

	return apiv1.StatusOK, respondents
}

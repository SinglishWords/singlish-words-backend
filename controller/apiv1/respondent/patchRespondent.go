package respondent

import (
	"github.com/gin-gonic/gin"
	"singlishwords/controller/apiv1"
	"singlishwords/service"
)

// PatchRespondent godoc
// @Summary Patch emails
// @Tags Respondent
// @Produce json
// @Param respondent body patchRespondentBody true "respondent information"
// @Success 204 ""
// @Failure 500 {object} apiv1.ErrorResponse
// @Router /respondent [patch]
func PatchRespondent(c *gin.Context) (apiv1.HttpStatus, interface{}) {
	code := apiv1.StatusOK
	rb := patchRespondentBody{}
	err := c.BindJSON(&rb)
	if err != nil {
		code = apiv1.StatusPostParamError
		return code, nil
	}

	err = service.UpdateRespondentEmail(rb.Id, rb.Email)
	if err != nil {
		return apiv1.StatusFail("Unable to update the email information."), nil
	}
	return apiv1.StatusNoContent, nil
}

type patchRespondentBody struct {
	Id    int64  `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
}

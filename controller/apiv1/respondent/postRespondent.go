package respondent

import (
	"singlishwords/controller/apiv1"
	"singlishwords/model"
	"singlishwords/service"
	"time"

	"github.com/gin-gonic/gin"
)

// PostRespondent godoc
// @Summary Post a respondent information
// @Tags Respondent
// @Produce json
// @Param respondent body postRespondentBody true "respondent information"
// @Success 201 {object} model.Respondent
// @Failure 500 {object} apiv1.ErrorResponse
// @Router /respondent [post]
func PostRespondent(c *gin.Context) (apiv1.HttpStatus, interface{}) {
	code := apiv1.StatusOK
	rb := postRespondentBody{}
	err := c.BindJSON(&rb)
	if err != nil {
		code = apiv1.StatusFail(err.Error()) //apiv1.StatusPostParamError
		return code, nil
	}

	r := rb.ToRespondent()
	err = service.AddRespondent(r)

	if err != nil {
		code = apiv1.StatusFail(err.Error())
		return code, nil
	}

	return apiv1.StatusCreated, r
}

type postRespondentBody struct {
	Age                int       `json:"age" db:"age"`
	Gender             string    `json:"gender" db:"gender"`
	Education          string    `json:"education" db:"education"`
	CountryOfBirth     string    `json:"countryOfBirth" db:"country_of_birth"`
	CountryOfResidence string    `json:"countryOfResidence" db:"country_of_residence"`
	Ethnicity          string    `json:"ethnicity" db:"ethnicity"`
	IsNative           string    `json:"isNative" db:"is_native"`
	LanguagesSpoken    string    `json:"languagesSpoken" db:"language_spoken"`
	StartTime          time.Time `json:"startTime" db:"start_time"`
	EndTime            time.Time `json:"endTime" db:"end_time"`
	Email              string    `json:"email" db:"email"`
}

func (rb *postRespondentBody) ToRespondent() *model.Respondent {
	return &model.Respondent{
		Age:                rb.Age,
		Gender:             rb.Gender,
		Education:          rb.Education,
		CountryOfBirth:     rb.CountryOfBirth,
		CountryOfResidence: rb.CountryOfResidence,
		Ethnicity:          rb.Ethnicity,
		IsNative:           rb.IsNative,
		LanguagesSpoken:    rb.LanguagesSpoken,
		StartTime:          rb.StartTime,
		EndTime:            rb.EndTime,
		Email:              rb.Email,
	}
}

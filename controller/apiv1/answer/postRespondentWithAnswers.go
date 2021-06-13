package answer

import (
	"github.com/gin-gonic/gin"
	"singlishwords/controller/apiv1"
	"singlishwords/model"
	"singlishwords/service"
	"time"
)

// PostRespondentWithAnswers godoc
// @Summary Post respondent information and answers together
// @Tags Answer
// @Produce json
// @Param respondent body paramPostRespondentWithAnswers true "the information"
// @Success 201 {object} model.Respondent
// @Failure 500 {object} apiv1.ErrorResponse
// @Router /answers [post]
func PostRespondentWithAnswers(c *gin.Context) (apiv1.HttpStatus, interface{}) {
	code := apiv1.StatusOK
	rb := paramPostRespondentWithAnswers{}
	err := c.BindJSON(&rb)
	if err != nil {
		code = apiv1.StatusPostParamError
		return code, nil
	}

	r, as := rb.ToModels()
	r, err = service.AddRespondentAndAnswersTogether(r, as)

	if err != nil {
		code = apiv1.StatusFail(err.Error())
		return code, nil
	}

	return apiv1.StatusCreated, r
}

type paramPostRespondentWithAnswers struct {
	Age                int       `json:"age"`
	Gender             string    `json:"gender"`
	Education          string    `json:"education"`
	CountryOfBirth     string    `json:"countryOfBirth"`
	CountryOfResidence string    `json:"countryOfResidence"`
	Ethnicity          string    `json:"ethnicity"`
	IsNative           string    `json:"isNative"`
	LanguagesSpoken    string    `json:"languagesSpoken"`
	StartTime          time.Time `json:"startTime"`
	EndTime            time.Time `json:"endTime"`
	Email              string    `json:"email"`
	Answers            []struct {
		QuestionId   int64  `json:"questionId"`
		TimeSpend    int    `json:"timeSpend"`
		Association1 string `json:"association1"`
		Association2 string `json:"association2"`
		Association3 string `json:"association3"`
	} `json:"answers"`
}

// ToModels
// Transform body parameter to respondent and answers as data model in db
func (rb *paramPostRespondentWithAnswers) ToModels() (*model.Respondent, []model.Answer) {
	r := &model.Respondent{
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
	answers := make([]model.Answer, len(rb.Answers))
	for i, answer := range rb.Answers {
		answers[i] = model.Answer{
			Association1: answer.Association1,
			Association2: answer.Association2,
			Association3: answer.Association3,
			TimeSpend:    time.Duration(answer.TimeSpend),
			QuestionId:   answer.QuestionId,
		}
	}
	return r, answers
}

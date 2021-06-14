package answer

import (
	"fmt"
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
		//code = apiv1.StatusPostParamError
		code = apiv1.StatusFail(err.Error())
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
	LanguagesSpoken    []string  `json:"languagesSpoken"`
	StartTime          time.Time `json:"startTime"`
	EndTime            time.Time `json:"endTime"`
	Email              string    `json:"email"`
	Answers            []struct {
		Question struct {
			Id   int64  `json:"id"`
			Word string `json:"word"`
		} `json:"question"`
		Responses [3]string `json:"response"`
		TimeSpend int       `json:"timeOnPage"`
	} `json:"data"`
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
		LanguagesSpoken:    fmt.Sprintf("%+q", rb.LanguagesSpoken),
		StartTime:          rb.StartTime,
		EndTime:            rb.EndTime,
		Email:              rb.Email,
	}
	answers := make([]model.Answer, len(rb.Answers))
	for i, answer := range rb.Answers {
		answers[i] = model.Answer{
			Association1: answer.Responses[0],
			Association2: answer.Responses[1],
			Association3: answer.Responses[2],
			TimeSpend:    time.Duration(answer.TimeSpend),
			QuestionId:   answer.Question.Id,
		}
	}
	return r, answers
}

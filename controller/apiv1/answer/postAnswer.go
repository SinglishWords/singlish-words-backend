package answer

import (
	"github.com/gin-gonic/gin"
	"singlishwords/controller/apiv1"
	"singlishwords/model"
	"singlishwords/service"
	"time"
)

type paramPostSingleAnswer struct {
	QuestionId   int64  `json:"questionId"`
	RespondentId int64  `json:"respondentId"`
	TimeSpend    int    `json:"timeSpend"`
	Association1 string `json:"association1"`
	Association2 string `json:"association2"`
	Association3 string `json:"association3"`
}

func (p *paramPostSingleAnswer) ToAnswer() *model.Answer {
	return &model.Answer{
		Id:           -1,
		Association1: p.Association1,
		Association2: p.Association2,
		Association3: p.Association3,
		TimeSpend:    time.Duration(p.TimeSpend),
		QuestionId:   p.QuestionId,
		RespondentId: p.RespondentId,
	}
}

// PostAnswer godoc
// @Summary Post an answer
// @Tags Answer
// @Produce json
// @Param answer body paramPostSingleAnswer true "answer with 3 associations"
// @Success 201 {object} model.Answer
// @Failure 500 {object} apiv1.ErrorResponse
// @Router /answer [post]
func PostAnswer(c *gin.Context) (apiv1.HttpStatus, interface{}) {
	var param paramPostSingleAnswer
	err := c.BindJSON(&param)

	if err != nil {
		return apiv1.StatusPostParamError, err
	}

	answer := param.ToAnswer()
	err = service.PostAnswer(answer)

	if err != nil {
		return apiv1.StatusFail(err.Error()), nil
	}

	return apiv1.StatusCreated, answer
}

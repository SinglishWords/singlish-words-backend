package answer

import (
	"singlishwords/controller/apiv1"
	"singlishwords/model"
	"singlishwords/service"
	"time"

	"github.com/gin-gonic/gin"
)

type paramPostSingleAnswer struct {
	QuestionId   int64     `json:"questionId"`
	RespondentId int64     `json:"respondentId"`
	TimeSpend    int       `json:"timeSpend"`
	Responses    [3]string `json:"response"`
}

func (p *paramPostSingleAnswer) ToAnswer() *model.Answer {
	return &model.Answer{
		Id:           -1,
		Association1: p.Responses[0],
		Association2: p.Responses[1],
		Association3: p.Responses[2],
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

	q, err := service.GetQuestionById(param.QuestionId)
	if err != nil {
		return apiv1.StatusFail(err.Error()), nil
	}

	for _, resp := range param.Responses {
		if resp != "" {
			err = service.IncrementAssociationCount(q.Word, resp, 1)
			if err != nil {
				return apiv1.StatusFail(err.Error()), nil
			}
		}			
	}

	return apiv1.StatusCreated, answer
}

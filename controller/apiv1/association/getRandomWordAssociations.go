package association

import (
	"fmt"
	"singlishwords/controller/apiv1"
	"singlishwords/log"
	"singlishwords/service"

	"github.com/gin-gonic/gin"
)

type RandomWordAssociations struct {
	Word string `json:"word"`
	Forward service.Visualisation `json:"forward"`
	Bacward service.Visualisation `json:"backward"`
}


func GetRandomWordAssociations(c *gin.Context) (apiv1.HttpStatus, interface{}) {
	randomQuestion, err := service.GetRandomQuestion()
	if err != nil {
		return apiv1.StatusFail(err.Error()), err
	}
	word := randomQuestion.Word

	log.Logger.Infof(fmt.Sprintf("Getting backward associations for word: %s", word))
	backwardAssociations, err := service.GetBackwardAssociations(word)
	if err != nil {
		return apiv1.StatusFail(err.Error()), err
	}

	log.Logger.Infof(fmt.Sprintf("Getting forward associations for word: %s", word))
	forwardAssociations, err := service.GetForwardAssociations(word)
	if err != nil {
		return apiv1.StatusFail(err.Error()), err
	}

	return apiv1.StatusOK, &RandomWordAssociations{Word: word, Forward: *forwardAssociations, Bacward: *backwardAssociations}
}

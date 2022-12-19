package association

import (
	"fmt"
	"singlishwords/controller/apiv1"
	"singlishwords/log"
	"singlishwords/service"

	"github.com/gin-gonic/gin"
)


func GetRandomWordBackwardAssociations(c *gin.Context) (apiv1.HttpStatus, interface{}) {
	randomQuestion, err := service.GetRandomQuestion()
	if err != nil {
		return apiv1.StatusFail(err.Error()), err
	}
	word := randomQuestion.Word
	log.Logger.Infof(fmt.Sprintf("Getting backward associations for word: %s", word))

	associations, err := service.GetBackwardAssociations(word)
	if err != nil {
		return apiv1.StatusFail(err.Error()), err
	}
	return apiv1.StatusOK, associations
}

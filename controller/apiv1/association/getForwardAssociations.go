package association

import (
	"fmt"
	"singlishwords/controller/apiv1"
	"singlishwords/log"
	"singlishwords/service"
	"strings"

	"github.com/gin-gonic/gin"
)


func GetForwardAssociations(c *gin.Context) (apiv1.HttpStatus, interface{}) {
	word := c.Param("word")
	word = strings.Replace(word, "-", " ", -1)
	log.Logger.Infof(fmt.Sprintf("Getting forward associations for word: %s", word))

	associations, err := service.GetForwardAssociations(word)
	if err != nil {
		return apiv1.StatusFail(err.Error()), err
	}
	return apiv1.StatusOK, associations
}

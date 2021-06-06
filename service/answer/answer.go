package anwser_service

import (
	"singlishwords/model"

	"github.com/gin-gonic/gin"
)

func PostAnswers() error {
	// TODO
	return nil
}

func GetAnswers(c *gin.Context) ([]model.Answer, error) {
	return []model.Answer{}, nil
}

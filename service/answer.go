package service

import (
	"singlishwords/dao"
	"singlishwords/model"
)

// import (
// 	"singlishwords/model/answer"

// 	"github.com/gin-gonic/gin"
// )

// func PostAnswers(respondentId int64, answers []answer.Answer) error {
// 	for i := range answers {
// 		answers[i].RespondentId = respondentId
// 	}

// 	answer.BulkSave(answers)
// 	return err
// }

var answerDAO = dao.AnswerDAO{}

func GetAnswers(limit int) ([]model.Answer, error) {
	return answerDAO.GetAll()
}

func PostAnswer(answer *model.Answer) error {
	return answerDAO.Save(answer)
}

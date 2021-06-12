package service

import (
	"singlishwords/dao"
	"singlishwords/model"
)

var answerDAO = dao.AnswerDAO{}

func GetAnswers(limit int) ([]model.Answer, error) {
	return answerDAO.GetAll()
}

func PostAnswer(answer *model.Answer) error {
	return answerDAO.Save(answer)
}

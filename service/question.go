package service

import (
	"singlishwords/cache"
	"singlishwords/dao"
	"singlishwords/model"
)

var questionCache cache.QuestionCache

func GetRandomNQuestions(limit int) ([]model.Question, error) {
	questions, err := questionCache.GetNRandomQuestions(limit)
	return questions, err
}

func AddQuestions() error {
	return nil
}

func DelQuestion() error {
	return nil
}

func GetAllQuestions() ([]model.Question, error) {
	questionDAO := dao.QuestionDAO{}
	questions, err := questionDAO.GetAll()
	if err != nil {
		return nil, err
	}
	return questions, nil
}

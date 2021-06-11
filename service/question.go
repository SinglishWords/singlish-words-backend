package service

import (
	"singlishwords/cache"
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

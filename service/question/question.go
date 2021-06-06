package question_service

import (
	"singlishwords/model"
)

func GetQuestions(limit int) ([]model.Question, error) {
	questions, err := CacheGetQuestions(limit)

	// questions := []model.Question{}
	// sql := "SELECT * FROM `question` ORDER BY RAND() LIMIT ?;"
	// err := db.Select(&questions, sql, limit)

	return questions, err
}

func AddQuestion() error {
	return nil
}

func DelQuestion() error {
	return nil
}

package dao

import (
	"singlishwords/model"
)

const (
	sqlSaveQuestion    = "INSERT INTO question (id, word) VALUES (:id, :word);"
	sqlGetQuestionById = "SELECT * FROM question WHERE id=?;"
	sqlGetAllQuestions = "SELECT * FROM question;"
)

type QuestionDAO struct{}

func (o QuestionDAO) GetAll() ([]model.Question, error) {
	var questions []model.Question
	err := db.Select(&questions, sqlGetAllQuestions)
	if err != nil {
		return questions, err
	}
	return questions, err
}

func (o QuestionDAO) Save(question *model.Question) error {
	result, err := db.NamedExec(sqlSaveQuestion, question)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	question.Id = id
	return err
}

func (o QuestionDAO) GetById(id int64) (model.Question, error) {
	var question model.Question
	err := db.Get(&question, sqlGetQuestionById, id)
	return question, err
}

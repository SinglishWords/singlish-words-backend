package dao

import (
	"singlishwords/model"
)

const (
	sqlGetAllAnswers = `SELECT * FROM answer;`
	sqlSaveAnswer    = `INSERT INTO answer 
						(association1, association2, association3, time_spend, question_id, respondent_id)
						VALUES 
						(:association1, :association2, :association3, :time_spend, :question_id, :respondent_id);`
)

type AnswerDAO struct{}

func (o AnswerDAO) GetAll() ([]model.Answer, error) {
	var answers []model.Answer
	err := db.Select(&answers, sqlGetAllAnswers)
	if err != nil {
		return answers, err
	}
	return answers, err
}

func (o AnswerDAO) Save(answer *model.Answer) error {
	result, err := db.NamedExec(sqlSaveAnswer, answer)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	answer.Id = id
	return err
}

func (o AnswerDAO) SaveAll(answers []model.Answer) error {
	_, err := db.NamedExec(sqlSaveAnswer, answers)
	return err
}

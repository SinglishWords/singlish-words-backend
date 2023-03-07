package dao

import (
	"singlishwords/database"
	"singlishwords/log"
	"singlishwords/model"
)

const (
	sqlGetAllAnswers = `SELECT * FROM answer;`
	sqlInsertAnswer  = `INSERT INTO answer 
						(association1, association2, association3, is_recognised_word, time_spend, question_id, respondent_id)
						VALUES 
						(:association1, :association2, :association3, :is_recognised_word, :time_spend, :question_id, :respondent_id);`
	sqlUpdateAnswer = `UPDATE answer SET association1=:association1, association2=:association2, association3=:association3 WHERE id=:id;`
)

type AnswerDAO struct{}

func (o AnswerDAO) GetAll() ([]model.Answer, error) {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}
	var answers []model.Answer
	err = db.Select(&answers, sqlGetAllAnswers)
	if err != nil {
		return answers, err
	}
	return answers, err
}

func (o AnswerDAO) Save(answer *model.Answer) error {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return notConnectedError{}
	}
	result, err := db.NamedExec(sqlInsertAnswer, answer)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	answer.Id = id
	log.Logger.Infof("Add a new answer to database: %+v", answer)
	return err
}

func (o AnswerDAO) SaveAll(answers []model.Answer) error {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return notConnectedError{}
	}
	_, err = db.NamedExec(sqlInsertAnswer, answers)
	log.Logger.Infof("Saved %d answers to database.", len(answers))
	return err
}

func (o AnswerDAO) UpdateAssociations(answer *model.Answer) error {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return notConnectedError{}
	}
	_, err = db.NamedExec(sqlUpdateAnswer, answer)
	if err != nil {
		return err
	}
	log.Logger.Infof("Updated answer %d", answer.Id)
	return err
}

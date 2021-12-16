package dao

import (
	"singlishwords/database"
	"singlishwords/log"
	"singlishwords/model"
)

const (
	sqlSaveQuestion    = "INSERT INTO question (id, word) VALUES (:id, :word);"
	sqlGetQuestionById = "SELECT * FROM question WHERE id=?;"
	sqlGetAllQuestions = "SELECT * FROM question WHERE enable=1;"
)

type QuestionDAO struct{}

func (o QuestionDAO) GetAll() ([]model.Question, error) {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}
	var questions []model.Question
	err = db.Select(&questions, sqlGetAllQuestions)
	if err != nil {
		log.Logger.Warn("Error when select all questions", err)
		return questions, err
	}
	return questions, err
}

func (o QuestionDAO) Save(question *model.Question) error {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return notConnectedError{}
	}
	result, err := db.NamedExec(sqlSaveQuestion, question)
	if err != nil {
		log.Logger.Warnw("Error when save a question",
			"err", err,
			"question", question)
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	question.Id = id
	log.Logger.Infof("Add a new question to database: %+v", question)
	return err
}

func (o QuestionDAO) GetById(id int64) (*model.Question, error) {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}
	var question model.Question
	err = db.Get(&question, sqlGetQuestionById, id)
	return &question, err
}

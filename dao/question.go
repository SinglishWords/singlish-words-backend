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
	sqlUpdateCount     = "UPDATE question SET count=count+1 WHERE id=?;"
	sqlMaxCount        = "SELECT MAX(count) FROM question"
	sqlGetWeighed      = "SELECT * FROM question ORDER BY -LOG(RAND()) / count DESC LIMIT ?;"
	sqlEntries         = "SELECT COUNT(id) FROM questions"
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

func (o QuestionDAO) UpdateCount(question *model.Question) error {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return notConnectedError{}
	}
	result, err := db.Exec(sqlUpdateCount, question.Id)
	if err != nil {
		log.Logger.Warnw("Error when updating count",
			"err", err,
			"question", question)
		return err
	} else if result != nil {
		log.Logger.Infof("Updated count for %d", question.Id)
	}
	return err
}

func (o QuestionDAO) GetMaxCount() (int64, error) {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return 0, notConnectedError{}
	}
	var maxi int64
	err = db.Get(&maxi, sqlMaxCount)
	return maxi + 1, err
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

func (o QuestionDAO) GetWeightedQuestions(limit int) ([]model.Question, error) {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}
	var questions []model.Question
	err = db.Select(&questions, sqlGetWeighed, limit)
	if err != nil {
		log.Logger.Warn("Error when select weighted questions", err)
		return questions, err
	}
	return questions, err
}

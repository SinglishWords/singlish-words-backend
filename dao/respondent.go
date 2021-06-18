package dao

import (
	"singlishwords/database"
	"singlishwords/log"
	"singlishwords/model"
)

const (
	sqlInsertRespondent = `INSERT INTO respondent 
			(age, gender, education, country_of_birth, country_of_residence, 
			 ethnicity, is_native, language_spoken, start_time, end_time, email, want_lucky_draw, want_update) 
			 VALUES (:age, :gender, :education, :country_of_birth, :country_of_residence, 
			 :ethnicity, :is_native, :language_spoken, :start_time, :end_time, :email, :want_lucky_draw, :want_lucky_draw);`
	sqlGetAllRespondents     = `SELECT * FROM respondent;`
	sqlGetRespondentById     = `SELECT * FROM respondent WHERE id=?;`
	sqlUpdateRespondentEmail = `UPDATE respondent SET email=?, want_lucky_draw=?, want_update=? WHERE id=?`
)

type RespondentDAO struct{}

func (RespondentDAO) Save(respondent *model.Respondent) error {
	db, err := database.GetMySqlDB()
	if db == nil {
		return notConnectedError{}
	}
	result, err := db.NamedExec(sqlInsertRespondent, respondent)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	respondent.Id = id
	return nil
}

func (RespondentDAO) GetById(id int64) (*model.Respondent, error) {
	db, err := database.GetMySqlDB()
	if db == nil {
		return nil, notConnectedError{}
	}
	var respondent model.Respondent
	err = db.Get(&respondent, sqlGetRespondentById, id)
	return &respondent, err
}

func (RespondentDAO) GetAll() ([]model.Respondent, error) {
	db, err := database.GetMySqlDB()
	if db == nil {
		return nil, notConnectedError{}
	}
	var respondents []model.Respondent
	err = db.Select(&respondents, sqlGetAllRespondents)
	return respondents, err
}

func (RespondentDAO) UpdateEmail(id int64, email string, wantLuckyDraw string, wantUpdate string) error {
	db, err := database.GetMySqlDB()
	if db == nil {
		return notConnectedError{}
	}

	_, err = db.Exec(sqlUpdateRespondentEmail, email, wantLuckyDraw, wantUpdate, id)
	if err != nil {
		log.Logger.Warnf("Fail to update a user's email, user id: %v, email: %v", id, email)
		return err
	}
	log.Logger.Infof("Successfully update a user's email, user id: %v, email: %v", id, email)
	return err
}

func (RespondentDAO) AddRespondentWithAnswers(r *model.Respondent, as []model.Answer) (*model.Respondent, error) {
	db, err := database.GetMySqlDB()
	if db == nil {
		return nil, notConnectedError{}
	}

	tx, err := db.Beginx()
	result, err := tx.NamedExec(sqlInsertRespondent, r)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Logger.Errorf("insert respondent failed: %v, unable to back: %v, the respondent: %+v", err, rollbackErr, r)
		}
		log.Logger.Warn(err)
		return nil, insertError{}
	}

	rid, err := result.LastInsertId()

	for _, a := range as {
		a.RespondentId = rid
		result, err = tx.NamedExec(sqlInsertAnswer, a)
		if err != nil {
			log.Logger.Warnf("insert answer failed: %v, the answer: %+v", err, a)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Logger.Warnf("insert respondent with answers together failed. The respondent: %v, the answers: %v", r, as)
		return nil, err
	}

	r.Id = rid
	log.Logger.Infof("Saved respondent with %d answers together. The respondent: %v, the answers: %v", len(as), r, as)
	return r, nil
}

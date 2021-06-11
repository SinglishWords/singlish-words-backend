package dao

import "singlishwords/model"

const (
	sqlInsertRespondent = `INSERT INTO respondent 
			(age, gender, education, country_of_birth, country_of_residence, 
			 ethnicity, is_native, language_spoken, start_time, end_time, email) 
			 VALUES (:age, :gender, :education, :country_of_birth, :country_of_residence, 
			 :ethnicity, :is_native, :language_spoken, :start_time, :end_time, :email);`
	sqlGetAllRespondents = `SELECT * FROM respondent;`
	sqlGetRespondentById = "SELECT * FROM respondent WHERE id=?;"
)

type RespondentDAO struct{}

func (RespondentDAO) Save(respondent *model.Respondent) error {
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

func (RespondentDAO) GetById(id int64) (model.Respondent, error) {
	var respondent model.Respondent
	err := db.Get(&respondent, sqlGetRespondentById, id)
	return respondent, err
}

func (RespondentDAO) GetAll() ([]model.Respondent, error) {
	var respondents []model.Respondent
	err := db.Select(&respondents, sqlGetAllRespondents)
	return respondents, err
}

package service

import (
	"singlishwords/dao"
	"singlishwords/model"
)

var respondentDAO dao.RespondentDAO

func AddRespondent(r *model.Respondent) error {
	err := respondentDAO.Save(r)
	return err
}

func GetAllRespondents() ([]model.Respondent, error) {
	return respondentDAO.GetAll()
}

func GetRespondentById(id int64) (*model.Respondent, error) {
	return respondentDAO.GetById(id)
}

func AddRespondentAndAnswersTogether(r *model.Respondent, as []model.Answer) (*model.Respondent, error) {
	return respondentDAO.AddRespondentWithAnswers(r, as)
}

func UpdateRespondentEmail(id int64, email string, wantLuckyDraw string, wantUpdate string) error {
	return respondentDAO.UpdateEmail(id, email, wantLuckyDraw, wantUpdate)
}

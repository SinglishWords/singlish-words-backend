package service

import (
	"singlishwords/dao"
	"singlishwords/model"
)

var emailDAO dao.EmailDAO

func AddEmail(e *model.Email) error {
	err := emailDAO.Save(e)
	return err
}

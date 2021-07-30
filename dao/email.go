package dao

import (
	"singlishwords/database"
	"singlishwords/log"
	"singlishwords/model"
)

const (
	sqlInsertEmail = `INSERT INTO email 
					(email, want_lucky_draw, want_update, time_on_pages)
					VALUES 
					(:email, :want_lucky_draw, :want_update, :time_on_pages);`
)

type EmailDAO struct{}

func (o EmailDAO) Save(email *model.Email) error {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return notConnectedError{}
	}
	_, err = db.NamedExec(sqlInsertEmail, email)
	if err != nil {
		return err
	}
	log.Logger.Infof("Add a new email to database: %+v", email)
	return err
}

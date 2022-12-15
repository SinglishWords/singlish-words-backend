package dao

import (
	"singlishwords/database"
	"singlishwords/log"
	"singlishwords/model"
)

const (
	sqlGetAssociationsByQid = `SELECT * FROM association WHERE question_id=?;`
	sqlGetAssociation = `SELECT * FROM association WHERE question_id=? AND association=?;`
	sqlInsertAssociation  = `INSERT INTO association 
						(question_id, association, count)
						VALUES 
						(:question_id, :association, :count);`
	sqlUpdateAssociation  = `UPDATE association SET
						count=?
						WHERE question_id=? AND association=?;`
)

type AssociationDAO struct{}

func (o AssociationDAO) GetAssociationsByQid(qid int64) ([]model.Association, error) {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}

	var associations []model.Association
	err = db.Select(&associations, sqlGetAssociationsByQid, qid)
	return associations, err
}


func (o AssociationDAO) GetAssociation(qid int64, associatedWord string) (*model.Association, error) {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}

	var association model.Association
	err = db.Get(&association, sqlGetAssociation, qid, associatedWord)
	return &association, err
}

func (o AssociationDAO) IncrementAssociationBy(qid int64, associatedWord string, inc int64) error {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return notConnectedError{}
	}

	var association model.Association
	err = db.Get(&association, sqlGetAssociation, qid, associatedWord)
	if err != nil {
		// If empty, create an entry that starts with count = 0
		association = model.Association{QuestionId: qid, Association: associatedWord, Count: 0}
		res, err := db.NamedExec(sqlInsertAssociation, association)
		if err != nil {
			return err
		}
		log.Logger.Infof("Created new association: %+v", res)
	}

	newCount := association.Count + inc

	_, err = db.Exec(sqlUpdateAssociation, newCount, qid, associatedWord)
	log.Logger.Infof("Incremented association count by: %d", inc)
	return err
}
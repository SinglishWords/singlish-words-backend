package dao

import (
	"singlishwords/database"
	"singlishwords/log"
	"singlishwords/model"
)

const (
	sqlGetAssociationsBySource = `SELECT * FROM association WHERE source=?;`
	sqlGetAssociation = `SELECT * FROM association WHERE source=? AND target=?;`
	sqlInsertAssociation  = `INSERT INTO association 
						(source, target, count)
						VALUES 
						(:source, :target, :count);`
	sqlUpdateAssociation  = `UPDATE association SET
						count=?
						WHERE source=? AND target=?;`
)

type AssociationDAO struct{}

func (o AssociationDAO) GetAssociationsByQid(q string) ([]model.Association, error) {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}

	var associations []model.Association
	err = db.Select(&associations, sqlGetAssociationsBySource, q)
	return associations, err
}


func (o AssociationDAO) GetAssociation(q, associatedWord string) (*model.Association, error) {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}

	var association model.Association
	err = db.Get(&association, sqlGetAssociation, q, associatedWord)
	return &association, err
}

func (o AssociationDAO) IncrementAssociationBy(q string, associatedWord string, inc int64) error {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return notConnectedError{}
	}

	var association model.Association
	err = db.Get(&association, sqlGetAssociation, q, associatedWord)
	if err != nil {
		// If empty, create an entry that starts with count = 0
		association = model.Association{Source: q, Target: associatedWord, Count: 0}
		res, err := db.NamedExec(sqlInsertAssociation, association)
		if err != nil {
			return err
		}
		log.Logger.Infof("Created new association: %+v", res)
	}

	newCount := association.Count + inc

	_, err = db.Exec(sqlUpdateAssociation, newCount, q, associatedWord)
	log.Logger.Infof("Incremented association count by: %d", inc)
	return err
}
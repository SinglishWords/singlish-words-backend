package dao

import (
	"fmt"
	"singlishwords/database"
	"singlishwords/log"
	"singlishwords/model"
	"strings"
)

const (
	sqlGetAssociationsBySource = `SELECT * FROM association WHERE source=? ORDER BY count DESC;`
	sqlGetBackwardAssociationsBySource = `SELECT * FROM association WHERE target=? ORDER BY count DESC;`
	sqlGetAssociation = `SELECT * FROM association WHERE source=? AND target=?;`
	sqlCountForwardAssociation = `SELECT COALESCE(SUM(count), 0) AS count FROM association WHERE target=?;`
	sqlCountBackwardAssociation = `SELECT COALESCE(SUM(count), 0) AS count FROM association WHERE source=?;`
	sqlInsertAssociation  = `INSERT INTO association 
						(source, target, count)
						VALUES 
						(:source, :target, :count);`
	sqlUpdateAssociation  = `UPDATE association SET
						count=?
						WHERE source=? AND target=?;`
)

type AssociationDAO struct{}

func (o AssociationDAO) GetAssociationsBySource(q string) ([]model.Association, error) {
	db, _ := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}

	var associations []model.Association
	err := db.Select(&associations, sqlGetAssociationsBySource, q)
	log.Logger.Infof("Executing GetAssociationsBySource: %+v", associations)
	return associations, err
}

func (o AssociationDAO) GetBackwardAssociationsBySource(q string) ([]model.Association, error) {
	db, _ := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}

	var associations []model.Association
	err := db.Select(&associations, sqlGetBackwardAssociationsBySource, q)
	log.Logger.Infof("Executing GetBackwardAssociationsBySource: %+v", associations)
	return associations, err
}


func (o AssociationDAO) GetAssociation(q, associatedWord string) (*model.Association, error) {
	db, _ := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}

	var association model.Association
	err := db.Get(&association, sqlGetAssociation, q, associatedWord)
	return &association, err
}

func joinWithQuotes(arr []string) string {
	var sb strings.Builder
	for i, s := range arr {
		sb.WriteString(fmt.Sprintf("\"%s\"", s))
		if i < len(arr)-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

func (o AssociationDAO) MultiSelectBySource(sources []string) ([]model.Association, error) {
	if len(sources) == 0 {
		log.Logger.Info("Slice length is 0, not executing MultiSelectBySource")
		return make([]model.Association, 0), nil
	}

	db, _ := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}

	var associations []model.Association
	err := db.Select(&associations, fmt.Sprintf("SELECT * FROM association WHERE source IN (%s) ORDER BY source, count DESC;", joinWithQuotes(sources)))
	log.Logger.Infof("Executing MultiSelectBySource: %+v", associations)
	return associations, err
}

func (o AssociationDAO) MultiSelectByTarget(sources []string) ([]model.Association, error) {
	if len(sources) == 0 {
		log.Logger.Info("Slice length is 0, not executing MultiSelectBySource")
		return make([]model.Association, 0), nil
	}

	db, _ := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}

	var associations []model.Association
	err := db.Select(&associations, fmt.Sprintf("SELECT * FROM association WHERE target IN (%s) ORDER BY target, count DESC;", joinWithQuotes(sources)))
	log.Logger.Infof("Executing MultiSelectByTarget: %+v", associations)
	return associations, err
}

func (o AssociationDAO) IncrementAssociationBy(q string, associatedWord string, inc int64) error {
	db, _ := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return notConnectedError{}
	}

	var association model.Association
	err := db.Get(&association, sqlGetAssociation, q, associatedWord)
	if err != nil {
		// If empty, create an entry that starts with count = 0
		association = model.Association{Source: q, Target: associatedWord, Count: 0}
		_, err := db.NamedExec(sqlInsertAssociation, association)
		if err != nil {
			return err
		}
		log.Logger.Infof("Created new association '%s' -> '%s'", q, associatedWord)
	}

	newCount := association.Count + inc

	_, err = db.Exec(sqlUpdateAssociation, newCount, q, associatedWord)
	log.Logger.Infof("Incremented association count of '%s' -> '%s' by: %d", q, associatedWord, inc)
	return err
}

type Count struct {
	N int64	`db:"count"`
}

// Result will be sorted by count descending
func (o AssociationDAO) CountForwardAssociations(words []string) ([]model.AssociationValue, error) {
	if len(words) == 0 {
		log.Logger.Info("Slice length is 0, not executing CountForwardAssociations")
		return make([]model.AssociationValue, 0), nil
	}

	db, _ := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}

	var associationValues []model.AssociationValue
	err := db.Select(&associationValues, 
		fmt.Sprintf(
			// No need coalesce since if target does not exist, it won't be included in the result
			"SELECT target AS word, SUM(count) AS count FROM association WHERE target in (%s) GROUP BY target ORDER BY count DESC;",
			joinWithQuotes(words),
		),
	)
	log.Logger.Infof("Executing CountForwardAssociations: %+v", associationValues)
	return associationValues, err
}

// Result will be sorted by count descending
func (o AssociationDAO) CountBackwardAssociations(words []string) ([]model.AssociationValue, error) {
	if len(words) == 0 {
		log.Logger.Info("Slice length is 0, not executing CountBackwardAssociations")
		return make([]model.AssociationValue, 0), nil
	}

	db, _ := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}

	var associationValues []model.AssociationValue
	err := db.Select(&associationValues, 
		fmt.Sprintf(
			// No need coalesce since if source does not exist, it won't be included in the result
			"SELECT source AS word, SUM(count) AS count FROM association WHERE source in (%s) GROUP BY source ORDER BY count DESC;",
			joinWithQuotes(words),
		),
	)
	log.Logger.Infof("Executing CountBackwardAssociations: %+v", associationValues)
	return associationValues, err
}

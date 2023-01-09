package dao

import (
	"database/sql"
	"fmt"
	"singlishwords/database"
	"singlishwords/log"
	"singlishwords/model"
)

const (
	sqlGetAllCommunityMappings = `SELECT community, word FROM community_map ORDER BY community;`
	sqlInsertCommunityMapping  = `INSERT INTO community_map
								(word, community)
								VALUES 
								(:word, :community);`
	sqlGetCommunityMappingByWord = `SELECT * FROM community_map WHERE word=?;`
	sqlUpdateCommunityMapping  = `UPDATE community_map SET
							community=?
							WHERE id=?;`
)

type CommunityDAO struct{}

func (o CommunityDAO) GetAll() ([]model.CommunityMapping, error) {
	db, _ := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}
	var cm []model.CommunityMapping
	err := db.Select(&cm, sqlGetAllCommunityMappings)
	if err != nil {
		return cm, err
	}
	return cm, err
}

func (o CommunityDAO) GetByWord(word string) (*model.CommunityMapping, error) {
	db, _ := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}
	var cm model.CommunityMapping
	err := db.Get(&cm, sqlGetCommunityMappingByWord, word)
	return &cm, err
}

func (o CommunityDAO) MultiSelectByWord(words []string) ([]model.CommunityMapping, error) {
	if len(words) == 0 {
		log.Logger.Info("Slice length is 0, not executing MultiSelectByWord")
		return make([]model.CommunityMapping, 0), nil
	}

	db, _ := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, notConnectedError{}
	}

	var mappings []model.CommunityMapping
	err := db.Select(&mappings, fmt.Sprintf("SELECT * FROM community_map WHERE word IN (%s);", joinWithQuotes(words)))
	log.Logger.Infof("Executing MultiSelectByWord: %+v", mappings)
	return mappings, err
}

func (o CommunityDAO) Upsert(cm *model.CommunityMapping) error {
	db, _ := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return notConnectedError{}
	}

	oldCm, err := o.GetByWord(cm.Word)

	// If not found, insert
	if err == sql.ErrNoRows {
		result, err := db.NamedExec(sqlInsertCommunityMapping, cm)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		cm.Id = id
		log.Logger.Infof("Added a new community mapping to database: %+v", cm)	
		return nil
	}
	
	// If found, update
	_, err = db.Exec(sqlUpdateCommunityMapping, cm.Community, oldCm.Id)
	cm.Id = oldCm.Id
	log.Logger.Infof("Updated community mapping to database: %+v", cm)
	
	return err
}

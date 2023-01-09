package main

import (
	"fmt"
	"singlishwords/database"
	"singlishwords/log"
)

const (
	sqlGetForwardAssociations = `SELECT q.id, q.word, a.association1, a.association2, a.association3 
									FROM question AS q JOIN answer AS a ON q.id=a.question_id 
									WHERE q.word=?;`
)

type forwardAssociationTriplet struct {
	Qid          int64         `db:"id"`
	Word         string        `db:"word"`
	Association1 string        `db:"association1"`
	Association2 string        `db:"association2"`
	Association3 string        `db:"association3"`
}

func getForwardAssociations(q string) ([]forwardAssociationTriplet, error) {
	db, err := database.GetMySqlDB()
	if db == nil {
		log.Logger.Error("Cannot connect to mysql database.")
		return nil, err
	}

	var associations []forwardAssociationTriplet
	err = db.Select(&associations, sqlGetForwardAssociations, q)
	return associations, err
}

func countAssociationsFrequencies(associations []forwardAssociationTriplet) map[string]int64 {
	count := make(map[string]int64)

	// FIX: Need to give ID to associations
	for _, triplets := range associations {
		as := []string{triplets.Association1, triplets.Association2, triplets.Association3}
		for _, a := range as {
			if a == "" {
				continue
			}
			_, ok := count[a]
			if !ok {
				count[a] = 0
			}
			count[a]++
		}
	}

	return count
}

func createAssociations() error {
	answers, err := answerDAO.GetAll()
	if err != nil {
		return err
	}
	
	
	for _, ans := range answers {
		q, err := questionDAO.GetById(ans.QuestionId)
		if err != nil {
			return err
		}
		associations, err := getForwardAssociations(q.Word)
		if err != nil {
			return err
		}
		counts := countAssociationsFrequencies(associations)

		fmt.Println("---START WORD---")
		fmt.Println("word:", associations[0].Word)
		for association, count := range counts {
			fmt.Println("------")
			fmt.Println("association:", association)
			fmt.Println("count:", count)
			fmt.Println("------")
			err = associationDAO.IncrementAssociationBy(q.Word, association, count)
			if err != nil {
				return err
			}
		}
		fmt.Println("---END WORD---")
	}

	return nil
}
package main

import (
	"fmt"
	"strconv"
)

func migrateAge() error {
	respodents, err := respondentDAO.GetAll()
	if err != nil {
		return err
	}
	
	for _, respondent := range respodents {
		age, err := strconv.Atoi(respondent.Age)
		if err != nil {
			continue
		}
		
		if age < 18 || age > 80 {
			fmt.Printf("Updating answer %d\n", respondent.Id)
			fmt.Printf("Before: %+v\n", respondent)

			if age < 18 {
				respondent.Age = "Less than 18"
			} else if age > 80 {
				respondent.Age = "More than 80"
			}

			fmt.Printf("After: %+v\n", respondent)
			err = respondentDAO.Update(&respondent)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
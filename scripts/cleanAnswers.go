package main

import (
	"fmt"
	"singlishwords/utils"
)

func cleanAnswers() error {
	answers, err := answerDAO.GetAll()
	if err != nil {
		return err
	}
	
	for i := range answers {
		ans := &answers[i]
		fmt.Printf("Updating answer %d\n", ans.Id)
		fmt.Printf("Before: %+v\n", ans)
		ans.Association1 = utils.CleanUpAnswer(ans.Association1)
		ans.Association2 = utils.CleanUpAnswer(ans.Association2)
		ans.Association3 = utils.CleanUpAnswer(ans.Association3)
		fmt.Printf("After: %+v\n", ans)
		err = answerDAO.UpdateAssociations(ans)
		if err != nil {
			return err
		}
	}

	return nil
}
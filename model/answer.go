package model

import (
	"time"
)

type Answer struct {
	Id           int64         `json:"id" db:"id"`
	Association1 string        `json:"association1" db:"association1"`
	Association2 string        `json:"association2" db:"association2"`
	Association3 string        `json:"association3" db:"association3"`
	TimeSpend    time.Duration `json:"timeSpend" db:"time_spend"`
	QuestionId   int64         `json:"questionId" db:"question_id"`
	RespondentId int64         `json:"respondentId" db:"respondent_id"`
}

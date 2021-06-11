package model

import (
	"time"
)

type Respondent struct {
	Id                 int64     `json:"id" db:"id"`
	Age                int       `json:"age" db:"age"`
	Gender             string    `json:"gender" db:"gender"`
	Education          string    `json:"education" db:"education"`
	CountryOfBirth     string    `json:"countryOfBirth" db:"country_of_birth"`
	CountryOfResidence string    `json:"countryOfResidence" db:"country_of_residence"`
	Ethnicity          string    `json:"ethnicity" db:"ethnicity"`
	IsNative           string    `json:"isNative" db:"is_native"`
	LanguagesSpoken    string    `json:"languagesSpoken" db:"language_spoken"`
	StartTime          time.Time `json:"startTime" db:"start_time"`
	EndTime            time.Time `json:"endTime" db:"end_time"`
	Email              string    `json:"email" db:"email"`
}

//func NewRespondent(id int, time string) *Respondent {
//	return &Respondent{id, time}
//}

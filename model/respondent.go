package model

import (
	"time"
)

type Respondent struct {
	Id                     int64     `json:"id" db:"id"`
	Age                    string    `json:"age" db:"age"`
	Gender                 string    `json:"gender" db:"gender"`
	Education              string    `json:"education" db:"education"`
	DurationOfSgpResidence string    `json:"durationOfSgpResidenceList" db:"duration_of_sgp_residence"`
	CountryOfBirth         string    `json:"countryOfBirth" db:"country_of_birth"`
	CountryOfResidence     string    `json:"countryOfResidence" db:"country_of_residence"`
	Ethnicity              string    `json:"ethnicity" db:"ethnicity"`
	IsNative               string    `json:"isNative" db:"is_native"`
	LanguagesSpoken        string    `json:"languagesSpoken" db:"language_spoken"`
	StartTime              time.Time `json:"startTime" db:"start_time"`
	EndTime                time.Time `json:"endTime" db:"end_time"`
	Email                  string    `json:"email" db:"email"`
	WantLuckyDraw          string    `json:"wantLuckyDraw" db:"want_lucky_draw"`
	WantUpdate             string    `json:"wantUpdate" db:"want_update"`
	UUID                   string    `json:"uuid" db:"uuid"`
}

//func NewRespondent(id int, time string) *Respondent {
//	return &Respondent{id, time}
//}

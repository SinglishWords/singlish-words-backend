package model

type Respondent struct {
	Id   int    `json:"id"`
	Time string `json:"time"`
}

func NewRespondent(id int, time string) *Respondent {
	return &Respondent{id, time}
}

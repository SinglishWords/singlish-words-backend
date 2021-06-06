package model

type Answer struct {
	Id           int    `json:"id" db:"id"`
	Word         string `json:"word" db:"word"`
	QuestionId   int    `json:"question_id" db:"question_id"`
	RespondentId int    `json:"respondent_id" db:"respondent_id"`
}

func NewAnswer(id int, word string, questionId int, respondentId int) *Answer {
	return &Answer{id, word, questionId, respondentId}
}

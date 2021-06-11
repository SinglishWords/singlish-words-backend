package model

import "encoding/json"

type Question struct {
	Id   int64  `json:"id" db:"id"`
	Word string `json:"word" db:"word"`
}

func NewQuestion(id int64, word string) *Question {
	return &Question{id, word}
}

func (q Question) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(q)
	return bytes, err
}

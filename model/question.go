package model

import "encoding/json"

type Question struct {
	Id   int64   `json:"id" db:"id"`
	Word string  `json:"word" db:"word"`
	Enable int64 `json:"enable" db:"enable"`
	Count int64  `json:"count" db:"count"`
}

func (q Question) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(q)
	return bytes, err
}

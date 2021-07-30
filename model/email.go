package model

import "encoding/json"

type Email struct {
	Email         string `json:"email" db:"email"`
	WantLuckyDraw string `json:"wantLuckyDraw" db:"want_lucky_draw"`
	WantUpdate    string `json:"wantUpdate" db:"want_update"`
}

func (e Email) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(e)
	return bytes, err
}

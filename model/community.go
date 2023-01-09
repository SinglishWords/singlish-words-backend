package model

type CommunityMapping struct {
	Id           	int64	`db:"id"`
	Word   			string	`db:"word"`
	Community 		int64	`db:"community"`
}

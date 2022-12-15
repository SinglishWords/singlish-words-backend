package model

type Association struct {
	Id           	int64	`db:"id"`
	QuestionId   	int64	`db:"question_id"`
	Association 	string	`db:"association"`
	Count 			int64	`db:"count"`
}

type Node struct {
	Id 				int64	`json:"id"`
	Name 			string	`json:"name"`
	SymbolSize 		int64 	`json:"symbolSize"`
	Value 			int64 	`json:"value"`
	Category 		int64 	`json:"category"`
}

type Link struct {
	Source int64 `json:"source"`
	Target int64 `json:"target"`
}

type Category struct {
	Name string `json:"name"`
}
package model

type Vote struct {
	PostId int `json:"posId"`
	UserId int `json:"userId"`
}

type VoteCount struct {
	PostId int `json:"posId"`
	Count  int `json:"count"`
}
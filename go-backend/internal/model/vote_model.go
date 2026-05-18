package model

type Vote struct {
	PostId int `json:"posId"`
	UserId int `json:"userId"`
}
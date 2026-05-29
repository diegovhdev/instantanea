package model

type Follow struct {
	FollowId    int `json:"followId"`
	FollowingId int `json:"followingId"`
}
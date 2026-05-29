package model

type Post struct {
	PostId   int    `json:"posId"`
	UserId   int    `json:"userId"`
	PublicId string `json:"publicID"`
	Url      string `json:"url"` 
	Text     string `json:"text"`
	Votes    int    `json:"votes"`
}

type PostResponse struct {
	PostId            int    `json:"postId"`
	UserId            int    `json:"userId"`
	Username          string `json:"username"`
	ProfilePictureUrl string `json:"profilePictureUrl"`
	Url               string `json:"url"` 
	Text              string `json:"text"`
	Votes             int    `json:"votes"`
	Voted             bool   `json:"voted"`  
	Following         bool   `json:"following"`
}
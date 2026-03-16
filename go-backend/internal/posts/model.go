package posts

type Post struct {
	PostId   int    `json:"posId"`
	UserId   int    `json:"userId"`
	PublicId string `json:"publicID"`
	Url      string `json:"url"` 
	Text     string `json:"text"`
}

type PostResponse struct {
	PostId   int    `json:"postId"`
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Url      string `json:"url"` 
	Text     string `json:"text"`
}
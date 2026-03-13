package posts

type Post struct {
	Id       int    `json:"id"`
	UserId   int    `json:"userId"`
	PublicId string `json:"publicID"`
	Url      string `json:"url"` 
	Text     string `json:"text"`
}


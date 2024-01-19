package models

type Evenisoide struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	Status      string `json:"status"`
	Created_at  string `json:"created_at"`
}

func (Evenisoide) TableName() string {
	return "blog_event"
}

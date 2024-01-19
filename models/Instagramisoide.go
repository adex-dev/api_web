package models

type Insta struct {
	Uid        int    `json:"uid"`
	Id         string `json:"id"`
	Permalink  string `json:"permalink"`
	Media_url  string `json:"media_url"`
	Thumbnail  string `json:"thumbnail"`
	Caption    string `json:"caption"`
	Media_type string `json:"media_type"`
	Timestamp  string `json:"timestamp"`
}

func (Insta) TableName() string {
	return "instagram_feed_isoide"
}

package models

type Promosi struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Addons      string `json:"addons"`
	Thumbnail   string `json:"thumbnail"`
	Status      string `json:"status"`
	Brands      string `json:"brands"`
}

func (Promosi) TableName() string {
	return "promotionblog"
}

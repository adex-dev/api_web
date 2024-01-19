package models

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Containt    string `json:"containt"`
	Categories  string `json:"categories"`
	Thumbnail   string `json:"thumbnail"`
	Brands      string `json:"brands"`
	Status      string `json:"status"`
}

func (Product) TableName() string {
	return "blog"
}

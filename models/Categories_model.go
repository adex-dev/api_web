package models

type Categories_tb struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Brand  string `json:"brand"`
	Status string `json:"status"`
}

func (Categories_tb) TableName() string {
	return "categories"
}

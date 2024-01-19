package models

type Store struct {
	Id_outlet   string `json:"id_outlet "`
	Nama_outlet string `json:"nama_outlet"`
	Thumbnail   string `json:"thumbnail"`
	Alamat      string `json:"alamat"`
	Add_ons     string `json:"add_ons"`
	Brands      string `json:"brands"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	Telp_1      string `json:"telp_1"`
	Telp_2      string `json:"telp_2"`
}

func (Store) TableName() string {
	return "prm_outlet"
}

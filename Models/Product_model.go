package models

type Product struct {
	Id         uint   `json:"id" gorm:"primaryKey"`
	NamaProduk string `json:"nama_produk" form:"nama_produk"`
	Jumlah     int    `json:"jumlah" form:"jumlah"`
	Harga      int    `json:"harga" form:"harga"`
	Publisher  string `json:"publisher" form:"publisher"`
	Image      string `json:"image" form:"image"`
	Deskripsi  string `json:"deskripsi" form:"deskripsi"`
}

func (Product) TableName() string {
	return "product"
}

package models

type ProductHistory struct {
	Id        uint `json:"id" gorm:"primaryKey"`
	UserId    uint `json:"user_id" form:"user_id"`
	ProductId uint `json:"product_id" form:"product_id"`
}

type DBGetProduct struct {
	Id     uint   `json:"id" gorm:"column:id"`
	UserId uint   `json:"user_id" gorm:"column:user_id"`
	Nama   string `json:"nama" gorm:"column:nama_produk"`
	Harga  int    `json:"harga" gorm:"column:harga"`
}

func (ProductHistory) TableName() string {
	return "product_history"
}

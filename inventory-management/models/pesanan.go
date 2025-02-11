package models

type Pesanan struct {
	ID_pesanan uint   `json:"id_pesanan" gorm:"primaryKey"`
	ID_produk  uint   `json:"id_produk"`
	Jumlah     int    `json:"jumlah"`
	Tanggal    string `json:"tanggal_pesanan"`
}

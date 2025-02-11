package models

type Inventaris struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ID_Produk uint   `json:"id_produk"`
	Jumlah    int    `json:"jumlah"`
	Lokasi    string `json:"lokasi"`
}

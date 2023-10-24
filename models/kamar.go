package models

type TipeKamar struct {
	ID          uint   `gorm:"primary_key" json:"id_tipe"`
	Description string `json:"description"`
	Fasilitas   string `json:"fasilitas"`
}

// Room represents the Rooms models with NamaKamar and a reference to TipeKamar.
type Kamar struct {
	ID          uint      `gorm:"primary_key" json:"id_kamar"`
	NamaKamar   string    `json:"namaKamar"`
	TipeKamarID uint      `json:"TipeKamarID"`
	TipeKamar   TipeKamar `gorm:"foreignKey:TipeKamarID" json:"TipeKamar"`
}

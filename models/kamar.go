package models

import "gorm.io/gorm"

type TipeKamar struct {
	gorm.Model
	ID          uint   `gorm:"primary_key" json:"IDTipeKamar"`
	Description string `json:"Deskripsi"`
	Fasilitas   string `json:"Fasilitas"`
}

// Room represents the Rooms models with NamaKamar and a reference to TipeKamar.
type Kamar struct {
	gorm.Model
	ID          uint      `gorm:"primary_key" json:"IDKamar"`
	NamaKamar   string    `json:"NamaKamar"`
	PhotoKamar  string    `json:"PhotoKamar"`
	TipeKamarID uint      `json:"IDTipeKamar"`
	TipeKamar   TipeKamar `gorm:"foreignKey:TipeKamarID" json:"TipeKamar"`
}

type KamarResponse struct {
	ID          uint
	Description string
	Fasilitas   string
}

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
	NamaKamar   string    `json:"NamaKamar" form:"NamaKamar"`
	PhotoKamar  string    `json:"PhotoKamar" form:"PhotoKamar"`
	TipeKamarID uint      `json:"IDTipeKamar" form:"IDTipeKamar"`
	TipeKamar   TipeKamar `gorm:"foreignKey:TipeKamarID" json:"TipeKamar"`
}

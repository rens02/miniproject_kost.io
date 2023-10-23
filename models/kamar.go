package models

import "gorm.io/gorm"

type TipeKamar struct {
	gorm.Model
	Description string `json:"description"`
	Fasilitas   string `json:"fasilitas"`
}

// Room represents the Rooms models with NamaKamar and a reference to TipeKamar.
type Kamar struct {
	gorm.Model
	NamaKamar   string    `json:"namaKamar"`
	TipeKamarID uint      `json:"TipeKamarID"`
	TipeKamar   TipeKamar `gorm:"foreignKey:TipeKamarID" json:"TipeKamar"`
}

// RoomAvailable represents the RoomAvailable models with waktu, status, price, and a reference to Rooms.
type KamarAvailable struct {
	gorm.Model
	Waktu   string  `json:"waktu"`
	Status  string  `json:"status"`
	Price   float64 `json:"price"`
	KamarID uint    `json:"KamarID"`
	Kamar   Kamar   `gorm:"foreignKey:KamarID json:"Kamar"`
}

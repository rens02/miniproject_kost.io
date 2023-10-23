package model

import "gorm.io/gorm"

type TipeKamar struct {
	gorm.Model
	Description string
}

// Room represents the Rooms model with NamaKamar and a reference to TipeKamar.
type Kamar struct {
	gorm.Model
	NamaKamar   string    `json:"namaKamar"`
	TipeKamarID uint      `json:"TipeKamarID"`
	TipeKamar   TipeKamar `gorm:"foreignKey:TipeKamarID" json:"TipeKamar"`
}

// RoomAvailable represents the RoomAvailable model with waktu, status, price, and a reference to Rooms.
type KamarAvailable struct {
	gorm.Model
	Waktu   string  `json:"waktu"`
	Status  string  `json:"status"`
	Price   float64 `json:"price"`
	KamarID uint    `json:"KamarID"`
	Kamar   Kamar   `gorm:"foreignKey:KamarID json:"Kamar"`
}

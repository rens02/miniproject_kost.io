package models

import "gorm.io/gorm"

type RoomStatus string

const (
	RoomStatusAvailable RoomStatus = "Available"
	RoomStatusOccupied  RoomStatus = "Kosong"
)

type Sewa struct {
	gorm.Model
	AvailableID   uint          `json:"RoomAvailableID"`
	UserID        uint          `json:"UserID"`
	KamarTersedia KamarTersedia `gorm:"foreignKey:AvailableID" json:"kamar_tersedia"`
	User          User          `gorm:"foreignKey:UserID" json:"user"`
}

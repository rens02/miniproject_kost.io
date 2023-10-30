package models

import "gorm.io/gorm"

type RoomStatus string
type RentStatus string

const (
	RoomStatusAvailable RoomStatus = "Available"
	RoomStatusOccupied  RoomStatus = "Occupied"
)
const (
	RentStatusBooked   RentStatus = "Booked"
	RentStatusCanceled RentStatus = "Canceled"
)

type Sewa struct {
	gorm.Model
	AvailableID   uint          `json:"RoomAvailableID"`
	UserID        uint          `json:"UserID"`
	RentStatus    RentStatus    `gorm:"type:enum('Booked', 'Canceled')" json:"RentStatus"`
	User          User          `gorm:"foreignKey:UserID" json:"User"`
	KamarTersedia KamarTersedia `gorm:"foreignKey:AvailableID" json:"KamarTersedia"`
}

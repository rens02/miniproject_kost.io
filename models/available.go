package models

type KamarTersedia struct {
	ID      uint       `gorm:"primary_key" json:"ID"`
	Waktu   string     `json:"Waktu"`
	Status  RoomStatus `gorm:"type:enum('Available', 'Occupied', 'Maintenance')" json:"Status"`
	Price   float64    `json:"Harga"`
	KamarID uint       `json:"IDKamar"`
	Kamar   Kamar      `gorm:"foreignKey:KamarID" json:"Kamar"`
}

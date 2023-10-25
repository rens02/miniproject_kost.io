package models

type KamarTersedia struct {
	ID      uint       `gorm:"primary_key" json:"id_available"`
	Waktu   string     `json:"waktu"`
	Status  RoomStatus `gorm:"type:enum('Available', 'Occupied', 'Maintenance')" json:"status"`
	Price   float64    `json:"harga"`
	KamarID uint       `json:"id_kamar"`
	Kamar   Kamar      `gorm:"foreignKey:KamarID" json:"Kamar"`
}

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

// RoomAvailable represents the RoomAvailable models with waktu, status, price, and a reference to Rooms.
type KamarTersedia struct {
	ID      uint       `gorm:"primary_key" json:"id_available"`
	Waktu   string     `json:"waktu"`
	Status  RoomStatus `gorm:"type:enum('Available', 'Occupied', 'Maintenance')" json:"status"`
	Price   float64    `json:"harga"`
	KamarID uint       `json:"id_kamar"`
	Kamar   Kamar      `gorm:"foreignKey:KamarID" json:"Kamar"`
}

// make a response struct without data redudancy
type KamarTersediaResponse struct {
	ID     uint    `json:"id_available"`
	Waktu  string  `json:"waktu"`
	Status string  `json:"status"`
	Price  float64 `json:"price"`
	Kamar  struct {
		NamaKamar string `json:"namaKamar"`
		TipeKamar struct {
			Description string `json:"description"`
			Fasilitas   string `json:"fasilitas"`
		} `json:"TipeKamar"`
	} `json:"Kamar"`
}

func ResponseAvail(kamarTersedia KamarTersedia) KamarTersediaResponse {
	response := KamarTersediaResponse{
		ID:     kamarTersedia.ID,
		Waktu:  kamarTersedia.Waktu,
		Status: string(kamarTersedia.Status),
		Price:  kamarTersedia.Price,
	}

	response.Kamar.NamaKamar = kamarTersedia.Kamar.NamaKamar
	response.Kamar.TipeKamar.Description = kamarTersedia.Kamar.TipeKamar.Description
	response.Kamar.TipeKamar.Fasilitas = kamarTersedia.Kamar.TipeKamar.Fasilitas

	return response
}

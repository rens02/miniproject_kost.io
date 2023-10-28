package helpers

import (
	"app/config"
	"app/models"
)

type KamarResponse struct {
	ID        uint   `json:"id_kamar"`
	NamaKamar string `json:"namaKamar"`
	TipeKamar struct {
		Description string `json:"description"`
		Fasilitas   string `json:"fasilitas"`
	} `json:"TipeKamar"`
}

type TipeKamarResponse struct {
	ID          uint   `json:"id_tipe"`
	Description string `json:"description"`
	Fasilitas   string `json:"fasilitas"`
}

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

func TipeKamarConvert(tipeKamar models.TipeKamar) TipeKamarResponse {
	response := TipeKamarResponse{
		ID:          tipeKamar.ID,
		Description: tipeKamar.Description,
		Fasilitas:   tipeKamar.Fasilitas,
	}

	return response
}

func KamarConvert(kamar models.Kamar) KamarResponse {
	response := KamarResponse{
		ID:        kamar.ID,
		NamaKamar: kamar.NamaKamar,
	}

	response.TipeKamar.Description = kamar.TipeKamar.Description
	response.TipeKamar.Fasilitas = kamar.TipeKamar.Fasilitas

	return response
}

func ResponseAvail(kamarTersedia models.KamarTersedia) KamarTersediaResponse {
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

func UpdateKamarTersediaAvailability(kamarTersediaID uint) error {

	// Find the KamarTersedia record by its ID
	var kamarTersedia models.KamarTersedia
	if err := config.DB.First(&kamarTersedia, kamarTersediaID).Error; err != nil {
		return err
	}

	// Update the availability status
	kamarTersedia.Status = models.RoomStatusOccupied

	// Save the updated record back to the database
	if err := config.DB.Save(&kamarTersedia).Error; err != nil {
		return err
	}

	return nil
}
func UpdateKamarTersediaAvailable(kamarTersediaID uint) error {

	// Find the KamarTersedia record by its ID
	var kamarTersedia models.KamarTersedia
	if err := config.DB.First(&kamarTersedia, kamarTersediaID).Error; err != nil {
		return err
	}

	// Update the availability status
	kamarTersedia.Status = models.RoomStatusAvailable

	// Save the updated record back to the database
	if err := config.DB.Save(&kamarTersedia).Error; err != nil {
		return err
	}

	return nil
}

type RentResponse struct {
	ID         int    `json:"RentID"`
	CreatedAt  string `json:"BookedAt"`
	RentStatus string `json:"RentStatus"`
	User       struct {
		Name  string `json:"Name"`
		Email string `json:"Email"`
	} `json:"user"`
	KamarTersedia struct {
		IDAvailable int    `json:"id_available"`
		Waktu       string `json:"waktu"`
		Status      string `json:"status"`
		Harga       int    `json:"harga"`
		Kamar       struct {
			NamaKamar string `json:"namaKamar"`
			TipeKamar struct {
				Description string `json:"description"`
				Fasilitas   string `json:"fasilitas"`
			} `json:"TipeKamar"`
		} `json:"Kamar"`
	} `json:"kamar_tersedia"`
}

func ResponseSewa(sewa models.Sewa) RentResponse {
	response := RentResponse{
		ID:         int(sewa.ID),
		CreatedAt:  sewa.CreatedAt.String(),
		RentStatus: string(sewa.RentStatus),
	}
	response.User.Name = sewa.User.Name
	response.User.Email = sewa.User.Email
	response.KamarTersedia.IDAvailable = int(sewa.KamarTersedia.ID)
	response.KamarTersedia.Waktu = sewa.KamarTersedia.Waktu
	response.KamarTersedia.Status = string(sewa.KamarTersedia.Status)
	response.KamarTersedia.Harga = int(sewa.KamarTersedia.Price)
	response.KamarTersedia.Kamar.NamaKamar = sewa.KamarTersedia.Kamar.NamaKamar
	response.KamarTersedia.Kamar.TipeKamar.Description = sewa.KamarTersedia.Kamar.TipeKamar.Description
	response.KamarTersedia.Kamar.TipeKamar.Fasilitas = sewa.KamarTersedia.Kamar.TipeKamar.Fasilitas

	return response
}

func ResponseHistory(sewa models.Sewa) RentResponse {
	response := RentResponse{
		ID:         int(sewa.ID),
		CreatedAt:  sewa.CreatedAt.String(),
		RentStatus: string(sewa.RentStatus),
	}
	response.User.Name = sewa.User.Name
	response.User.Email = sewa.User.Email
	response.KamarTersedia.IDAvailable = int(sewa.KamarTersedia.ID)
	response.KamarTersedia.Waktu = sewa.KamarTersedia.Waktu
	response.KamarTersedia.Status = "Canceled"
	response.KamarTersedia.Harga = int(sewa.KamarTersedia.Price)
	response.KamarTersedia.Kamar.NamaKamar = sewa.KamarTersedia.Kamar.NamaKamar
	response.KamarTersedia.Kamar.TipeKamar.Description = sewa.KamarTersedia.Kamar.TipeKamar.Description
	response.KamarTersedia.Kamar.TipeKamar.Fasilitas = sewa.KamarTersedia.Kamar.TipeKamar.Fasilitas

	return response
}

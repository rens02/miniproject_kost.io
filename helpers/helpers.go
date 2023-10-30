package helpers

import (
	"app/config"
	"app/models"
)

type KamarResponse struct {
	ID         uint   `json:"IDKamar"`
	NamaKamar  string `json:"NamaKamar"`
	PhotoKamar string `json:"PhotoKamar"`
	TipeKamar  struct {
		Description string `json:"DeskripsiKamar"`
		Fasilitas   string `json:"Fasilitas"`
	} `json:"TipeKamar"`
}

type TipeKamarResponse struct {
	ID          uint   `json:"IDTipeKamar"`
	Description string `json:"DeskripsiKamar"`
	Fasilitas   string `json:"Fasilitas"`
}

type KamarTersediaResponse struct {
	ID     uint    `json:"IDKamarTersedia"`
	Waktu  string  `json:"Waktu"`
	Status string  `json:"Status"`
	Price  float64 `json:"Harga"`
	Kamar  struct {
		NamaKamar string `json:"NamaKamar"`
		TipeKamar struct {
			Description string `json:"DeskripsiKamar"`
			Fasilitas   string `json:"FasilitasKamar"`
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
		ID:         kamar.ID,
		NamaKamar:  kamar.NamaKamar,
		PhotoKamar: kamar.PhotoKamar,
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
	ID         int    `json:"IDSewa"`
	CreatedAt  string `json:"DipesanPada"`
	RentStatus string `json:"StatusSewa"`
	User       struct {
		Name  string `json:"Nama"`
		Email string `json:"Email"`
	} `json:"User"`
	KamarTersedia struct {
		IDAvailable int    `json:"IDKamarTersedia"`
		Waktu       string `json:"Waktu"`
		Status      string `json:"Status"`
		Harga       int    `json:"Harga"`
		Kamar       struct {
			NamaKamar string `json:"NamaKamar"`
			TipeKamar struct {
				Description string `json:"DeskripsiKamar"`
				Fasilitas   string `json:"FasilitasKamar"`
			} `json:"TipeKamar"`
		} `json:"Kamar"`
	} `json:"KamarTersedia"`
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

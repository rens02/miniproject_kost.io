package helpers

import "app/models"

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

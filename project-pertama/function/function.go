package function

import "project-pertama/models"

func DataPeserta() []models.Peserta {
	//nama function harus kapital utk dapat dipanggil di class lainnya
	pesertaList := []models.Peserta{
		{
			Nama:      "Puji",
			Alamat:    "Jl kenangan",
			Pekerjaan: "Coder",
			Alasan:    "None",
		},
		{
			Nama:      "Soimah",
			Alamat:    "Jl Soto",
			Pekerjaan: "Host",
			Alasan:    "None",
		},
		{
			Nama:      "Jojo",
			Alamat:    "Jl Kita",
			Pekerjaan: "Tukang Pukul",
			Alasan:    "None",
		},
	}
	return pesertaList
}

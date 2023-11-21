package method

import (
	"fmt"
	"project-pertama/function"
	"project-pertama/models"
	"strings"
)

type CustomPeserta struct {
	models.Peserta
}

func (p CustomPeserta) FindPeserta() {
	var pesertas = function.DataPeserta()

	for key := range pesertas {
		if strings.ToLower(pesertas[key].Nama) == strings.ToLower(p.Nama) {
			fmt.Println("ID:", key+1)
			fmt.Println("Nama:", pesertas[key].Nama)
			fmt.Println("Alamat:", pesertas[key].Alamat)
			fmt.Println("Pekerjaan:", pesertas[key].Pekerjaan)
			fmt.Println("Alasan:", pesertas[key].Alasan)
		}
	}
}

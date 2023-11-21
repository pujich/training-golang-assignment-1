package main

import (
	"fmt"
	"os"
	"project-pertama/method"
	"project-pertama/models"
)

func main() {

	var args = os.Args

	if len(args) < 2 {
		fmt.Println("Usage: main <nama>")
		return
	}

	nama := args[1]
	//input nama peserta
	pesrta := method.CustomPeserta{
		Peserta: models.Peserta{
			Nama: nama,
		},
	}
	pesrta.FindPeserta()

}

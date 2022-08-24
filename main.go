package main

import (
	"fmt"
	"os"
	"strconv"
)

type data struct {
	name      string
	pekerjaan string
	alamat    string
	alasan    string
}

func main() {
	structAssignment()
}

func structAssignment() {

	list := []data{
		{
			name:      "Fadli Ramadhan",
			pekerjaan: "Programmer",
			alamat:    "Bekasi",
			alasan:    "Agar menjadi billgates",
		},
		{
			name:      "Manda",
			pekerjaan: "Cheft",
			alamat:    "Jakarta",
			alasan:    "Belajar Pemrograman",
		},
		{
			name:      "Hanif",
			pekerjaan: "Direktur",
			alamat:    "Depok",
			alasan:    "Menjadi Saingan Jackma",
		},
	}

	args := os.Args
	indexArgs := args[1]
	index, _ := strconv.Atoi(indexArgs)

	for i, t := range list {
		if i == index-1 {
			fmt.Println(t)
		}
	}

}

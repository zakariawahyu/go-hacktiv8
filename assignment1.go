package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Nama      string
	Alamat    string
	Pekerjaan string
}

func getStruct(number int) Person {
	person := []Person{
		{Nama: "Zakaria", Alamat: "Jakarta Pusat", Pekerjaan: "Backend"},
		{Nama: "Wahyu", Alamat: "Jakarta Barat", Pekerjaan: "Developer"},
		{Nama: "Nur", Alamat: "Jakarta Utara", Pekerjaan: "Developer"},
	}
	return person[number]
}

func main() {
	args := os.Args[1]
	i, _ := strconv.Atoi(args)
	fmt.Println(getStruct(i))
}

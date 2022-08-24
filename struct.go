package main

import "fmt"

type People struct {
	nama string
}

func main() {
	var closureFunction = func(names []*People) []string {
		var name []string
		for _, val := range names {
			name = append(name, val.nama)
		}
		return name
	}

	people := []*People{
		{nama: "Zakaria"},
		{nama: "Wahyu"},
		{nama: "Nur"},
		{nama: "Utomo"},
		{nama: "Fajar"},
		{nama: "Agus"},
		{nama: "Maulana"},
		{nama: "Rizki"},
		{nama: "Ramadan"},
	}

	fmt.Println(closureFunction(people))
}

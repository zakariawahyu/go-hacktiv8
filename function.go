package main

import (
	"fmt"
)

func main() {
	people := make([]string, 0, 10)
	people = append(people, "zaka", "fajar", "zayan", "rizki", "hendar", "billy", "maulana", "agus", "ramadan", "utomo")
	loopData(people)
}

func loopData(people []string) {
	for _, val := range people {
		cetak(val)
	}
}

func cetak(name string) {
	fmt.Println(name)
}

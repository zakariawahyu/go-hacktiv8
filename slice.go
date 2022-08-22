package main

import "fmt"

func main() {
	people := make([]string, 10)
	people = append(people, "zaka", "fajar", "zayan", "rizki", "hendar", "billy", "maulana", "agus", "ramadan", "utomo")

	for _, val := range people {
		fmt.Println(val)
	}
}

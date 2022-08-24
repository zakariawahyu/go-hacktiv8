package main

import "fmt"

func main() {
	var closureFunction = func(names []*string) []string {
		var name []string
		for _, val := range names {
			name = append(name, *val)
		}
		return name
	}

	zaka := "zakaria"
	wahyu := "wahyu"
	nur := "nur"
	utomo := "utomo"
	rizki := "rizki"
	ramadhan := "ramadhan"
	fajar := "fajar"
	agus := "agus"
	maulana := "maulana"

	people := []*string{&zaka, &wahyu, &nur, &utomo, &rizki, &ramadhan, &fajar, &agus, &maulana}
	fmt.Println(closureFunction(people))
}

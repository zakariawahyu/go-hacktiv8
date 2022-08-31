package main

import (
	"fmt"
	"sync"
)

func testAsyn(group *sync.WaitGroup, number int) {
	defer group.Done()

	fmt.Println("Ini number ke : ", number)
}

func main() {
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		group.Add(1)
		go testAsyn(group, i)
	}

	group.Wait()
	fmt.Println("Done")
}

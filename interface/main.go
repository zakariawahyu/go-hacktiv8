package main

import (
	"fmt"
	"go-hacktiv8/interface/service"
)

func main() {
	userServices := service.NewUserServices()
	result := userServices.Register(&service.User{
		"Zaka",
	})
	fmt.Println(result)
}

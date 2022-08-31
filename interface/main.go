package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-hacktiv8/interface/service"
)

func main() {
	var db []*service.User
	userServices := service.NewUserServices(db)

	//mux := http.NewServeMux()
	//mux.HandleFunc("/post-user", userServices.PostUserServer)
	//mux.HandleFunc("/get-user", userServices.GetUserServer)
	//
	//http.ListenAndServe("localhost:8081", mux)

	router := gin.Default()

	router.GET("/user", userServices.GetUserGin)
	router.POST("/user", userServices.PostUserGin)
	router.GET("/user/:id", userServices.GetUserByIdGin)
	router.Run("localhost:8081")
}

package main

import (
	"github.com/zakariawahyu/go-hacktiv8/interface/service"
	"net/http"
)

func main() {
	var db []*service.User
	userServices := service.NewUserServices(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/post-user", userServices.PostUserServer)
	mux.HandleFunc("/get-user", userServices.GetUserServer)

	http.ListenAndServe("localhost:8081", mux)
}

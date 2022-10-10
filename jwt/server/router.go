package server

import (
	"net/http"
)

func Start() {
	mux := http.NewServeMux()

	//endpoint := http.HandlerFunc(controller.GetUser)

	//mux.HandleFunc("/users", Midlleware1(Middleware2()))

	http.ListenAndServe("localhost:8081", mux)
}

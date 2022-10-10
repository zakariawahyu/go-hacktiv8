package server

import (
	"fmt"
	"net/http"
)

func Middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//before
		fmt.Println("Middleware 1 : Before")
		next.ServeHTTP(w, r)

		//after
		fmt.Println("Middleware 1 : After")
	})
}

func Middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware 2 : Before")

		next.ServeHTTP(w, r)

		fmt.Println("Middleware 3 : After")
	})
}

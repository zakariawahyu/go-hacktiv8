package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get User")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "hello world",
	})
}

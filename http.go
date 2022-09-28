package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	userID int
	ID     int
	Title  string
	Body   string
}

var url = "https://jsonplaceholder.typicode.com/posts"
var port = ":8081"

func main() {
	//HttpGet(url + "/1")
	//HttpPost(url)

	http.HandleFunc("/posts", AllowOnlyGet(Auth(GetPostById)))
	http.HandleFunc("/posts/create", AllowOnlyPost(Auth(CreatePost)))

	log.Println("Server running on port : ", port)
	http.ListenAndServe(port, nil)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	resp, err := HttpGet(url + "/" + query.Get("id"))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error,
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"payload": resp,
	})
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payload Post
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error,
		})
		return
	}

	resp, err := HttpPost(url, &payload)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"payload": resp,
	})
}

func AllowOnlyGet(next http.HandlerFunc) http.HandlerFunc {
	//return method == http.MethodGet
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Write([]byte("Only Get"))
		}
		next(w, r)
	}
}

func AllowOnlyPost(next http.HandlerFunc) http.HandlerFunc {
	//return method == http.MethodGet
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Write([]byte("Only Get"))
		}
		next(w, r)
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "need auth",
			})
		}

		if username != "admin" && password != "admin" {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "username / password salah",
			})
		}
	}
}

func get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	var data interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", data)
}

func HttpGet(url string) (interface{}, error) {
	data, err := req(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(data)

	return data, nil
}

func HttpPost(url string, reqPayload *Post) (interface{}, error) {
	data, err := json.Marshal(reqPayload)
	if err != nil {
		return nil, err
	}

	resp, err := req(http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)

	return resp, nil
}

func req(method, url string, body []byte) (map[string]interface{}, error) {
	client := http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

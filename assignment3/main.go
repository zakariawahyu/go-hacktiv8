package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"path"
	"text/template"
)

type DataJson struct {
	Water       int    `json:"water"`
	WaterStatus string `json:"water_status"`
	Wind        int    `json:"wind"`
	WindStatus  string `json:"wind_status"`
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8081", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c := make(chan DataJson)
	go UpdateData(c)
	data := <-c

	switch {
	case data.Water <= 5:
		data.WaterStatus = "Aman"
	case data.Water >= 6 && data.Water <= 8:
		data.WaterStatus = "Siaga"
	case data.Water > 8:
		data.WaterStatus = "Bahaya"
	default:
		break
	}

	switch {
	case data.Wind <= 6:
		data.WindStatus = "Aman"
	case data.Wind >= 7 && data.Wind <= 15:
		data.WindStatus = "Siaga"
	case data.Wind > 15:
		data.WindStatus = "Bahaya"
	default:
		break
	}

	var pathHtml = path.Join("assignment3/views/index.html")
	tmpl, err := template.ParseFiles(pathHtml)
	if err != nil {
		log.Printf("Error open file with err: %s", err)
	}
	newJson, _ := json.Marshal(data)
	ioutil.WriteFile("assignment3/views/data.json", newJson, 0655)
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateData(c chan DataJson) {
	for {
		data := DataJson{}

		randomWater := rand.Intn(15)
		randomWind := rand.Intn(20)
		data.Water = randomWater
		data.Wind = randomWind
		c <- data
	}
}

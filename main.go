package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// args := os.Args
	api_key := os.Getenv("OPENWEATHER_API_KEY")
	city := "cairo"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%v&appid=%v", city, api_key)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	w := Weather{}
	decoder := json.NewDecoder(resp.Body)
	// decoder.DisallowUnknownFields()

	if err := decoder.Decode(&w); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", w.Coord)
}

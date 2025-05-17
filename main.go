package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("please provid a city name as an argument to the command")
		os.Exit(1)
	}

	api_key := os.Getenv("OPENWEATHER_API_KEY")
	if api_key == "" {
		log.Fatal("environment variable OPENWEATHER_API_KEY is not set")
	}

	city := strings.Join(args[1:], "%20")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%v&appid=%v&units=metric", city, api_key)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status + ", the city name maybe incorrect")
	}

	w := Weather{}
	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&w); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current temprutere in %s is %2.0f with %s\n", w.Name, w.Main.Temp, w.Weather[0].Description)
}

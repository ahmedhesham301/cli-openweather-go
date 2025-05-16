package main

import (
	"fmt"
	"io"
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
	io.Copy(os.Stdout, resp.Body)
}

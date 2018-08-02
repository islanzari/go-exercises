package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type dataIP struct {
	IPAddress string  `json:"query"`
	HostName  string  `json:"org"`
	City      string  `json:"city"`
	Region    string  `json:"region"`
	Coutry    string  `json:"country"`
	Lon       float64 `json:"lon"`
	Lat       float64 `json:"lat"`
	Postal    string  `json:"zip"`
}

func main() {
	var ip string

	flag.StringVar(&ip, "ip", "8.8.8.9", "a string var")
	flag.Parse()

	req, err := http.NewRequest(http.MethodGet, "http://ip-api.com/json/"+ip, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var data dataIP
	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("IP address: ", data.IPAddress)
	fmt.Println("Hostname: ", data.HostName)
	fmt.Println("City: ", data.City)
	fmt.Println("Region: ", data.Region)
	fmt.Println("Country: ", data.Coutry)
	fmt.Println("Loc: ", data.Lon, " ", data.Lat)
	fmt.Println("Postal: ", data.Postal)

}

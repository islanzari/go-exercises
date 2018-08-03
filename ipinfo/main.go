package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/islanzari/go-exercises/ipinfo/request"
	"github.com/islanzari/go-exercises/ipinfo/valid"
)

func writingOut(data request.DataIP) {
	fmt.Println("IP address: ", data.IPAddress)
	fmt.Println("Organization: ", data.HostName)
	fmt.Println("City: ", data.City)
	fmt.Println("Region: ", data.Region)
	fmt.Println("Country: ", data.Coutry)
	fmt.Println("Loc: ", data.Lon, " ", data.Lat)
	fmt.Println("Postal: ", data.Postal)
}
func Geo(data request.DataIP) {
	fmt.Println("Loc: ", data.Lon, " ", data.Lat)
	fmt.Println("City: ", data.City)
	fmt.Println("Region: ", data.Region)
}

func main() {
	var ip string
	var geo bool
	flag.StringVar(&ip, "ip", "8.8.8.9", "a string var")
	flag.BoolVar(&geo, "geo", false, "a bool var")
	flag.Parse()

	if !valid.ValidIP4(ip) {
		fmt.Println("IP is invalids")
		return
	}

	data, err := request.RequestIP(ip)
	if err != nil {
		log.Println(err)
		return
	}
	if geo == true {
		Geo(data)
	} else {
		writingOut(data)
	}
}

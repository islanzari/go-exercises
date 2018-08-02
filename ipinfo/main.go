package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
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

func verificationIP(ip string) error {
	if !validIP4(ip) {
		err := errors.New("zle podaj prawidlowe ip")
		return err
	}
	return nil
}

func validIP4(ipAddress string) bool {
	ipAddress = strings.Trim(ipAddress, " ")
	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ipAddress) {
		return true
	}
	return false
}

func requestIP(ip string) (dataIP, error) {
	var data dataIP
	req, err := http.NewRequest(http.MethodGet, "http://ip-api.com/json/"+ip, nil)
	if err != nil {
		return data, err

	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return data, err
	}
	return data, err
}

func writingOut(data dataIP) {
	fmt.Println("IP address: ", data.IPAddress)
	fmt.Println("Organization: ", data.HostName)
	fmt.Println("City: ", data.City)
	fmt.Println("Region: ", data.Region)
	fmt.Println("Country: ", data.Coutry)
	fmt.Println("Loc: ", data.Lon, " ", data.Lat)
	fmt.Println("Postal: ", data.Postal)
}

func main() {
	var ip string
	flag.StringVar(&ip, "ip", "8.8.8.9", "a string var")
	flag.Parse()

	err := verificationIP(ip)
	if err != nil {
		log.Println(err)
		return
	}

	data, err := requestIP(ip)
	if err != nil {
		log.Println(err)
		return
	}
	writingOut(data)
}

package request

import (
	"encoding/json"
	"net/http"
)

type DataIP struct {
	IPAddress string  `json:"query"`
	HostName  string  `json:"org"`
	City      string  `json:"city"`
	Region    string  `json:"region"`
	Coutry    string  `json:"country"`
	Lon       float64 `json:"lon"`
	Lat       float64 `json:"lat"`
	Postal    string  `json:"zip"`
}

func RequestIP(a string) (DataIP, error) {
	var data DataIP
	req, err := http.NewRequest(http.MethodGet, "http://ip-api.com/json/"+a, nil)
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
	return data, err
}

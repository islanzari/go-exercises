package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func GetUser() error {
	var id string
	var rData struct {
		Status string `json:"status"`
		Data   struct {
			ID      uint64 `json:"id"`
			Name    string `json:"name"`
			Surname string `json:"surname"`
			Email   string `json:"email"`
		} `json:"data"`
	}

	client := &http.Client{}
	fmt.Println("podaj ID")
	fmt.Scan(&id)
	req, err := http.NewRequest("GET", *adres+"/users/"+id+"/", nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("error while fetching user")
	}

	err = json.NewDecoder(resp.Body).Decode(&rData)
	if err != nil {
		return err
	}
	log.Println(rData)
	return nil
}

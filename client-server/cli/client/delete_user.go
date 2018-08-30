package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func DeleteUser() error {
	var rData struct {
		Status string `json:"status"`
	}
	var id string
	client := &http.Client{}
	fmt.Println("podaj ID: ")
	fmt.Scan(&id)

	req, err := http.NewRequest("DELETE", *adres+"/users/"+id+"/", nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(&rData)
	if err != nil {
		return err
	}
	log.Println(rData)
	return nil
}

package client

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

var adres = flag.String("port", "http://localhost:8080", " server addres")

func Login() error {
	var rData struct {
		Status string `json:"status"`
		Data   struct {
			ID      uint64 `json:"id"`
			Name    string `json:"name"`
			Surname string `json:"surname"`
			Email   string `json:"email"`
		} `json:"data"`
	}
	b, err := bufer(form())
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", *adres+"/users/", b)
	if err != nil {
		return err
	}

	client := &http.Client{}
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

func bufer(payload Data) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	err := enc.Encode(payload)
	return buf, err
}

func form() Data {
	var name, surname, email string
	fmt.Println("wpisz imie")
	fmt.Scan(&name)
	fmt.Println("wpisz nazwisko")
	fmt.Scan(&surname)
	fmt.Println("wpisz email")
	fmt.Scan(&email)
	payload := Data{Name: name, Surname: surname, Email: email}
	return payload
}

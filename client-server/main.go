package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/islanzari/go-exercises/client-server/controler/users"
	"github.com/islanzari/go-exercises/client-server/model"
)

var adres = flag.String("port", ":8080", " application adress")
var tSave = flag.Int("time", 10000000000, "time save")
var db = flag.String("db", "", "database file")

func main() {
	flag.Parse()
	r := mux.NewRouter()
	usersModel := model.New(time.Duration(*tSave))
	if *db != "" {
		usersModel.LoadModel(*db)
	}
	usersHandler := users.Handle{
		Users: usersModel,
	}

	r.HandleFunc("/users/", usersHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}/", usersHandler.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}/", usersHandler.GetUser).Methods("GET")

	s := &http.Server{
		Addr:    *adres,
		Handler: r,
	}
	log.Fatal(s.ListenAndServe())
}

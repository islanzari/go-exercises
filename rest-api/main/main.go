package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/islanzari/go-exercises/rest-api/controler"
	"github.com/islanzari/go-exercises/rest-api/model"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable")
	if err != nil {
		panic(err)
	}
	todosModel := model.TodosModel{
		DB: db,
	}
	router := httprouter.New()
	handle := controler.Handle{
		Model: todosModel,
	}
	router.POST("/api/todos/", handle.CreateToDo)
	router.GET("/api/todos/", handle.FetchToDos)
	router.GET("/api/todos/:id/", handle.FetchToDo)
	router.PATCH("/api/todos/:id/", handle.UpdateToDo)
	router.DELETE("/api/todos/:id/", handle.DeleteToDo)
	log.Fatal(http.ListenAndServe(":8080", router))
}

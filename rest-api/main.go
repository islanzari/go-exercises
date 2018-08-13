package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/islanzari/go-exercises/rest-api/model"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

type Handle struct {
	Model model.TodosModel
}

func (h *Handle) DeleteToDo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	i, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Println("error bad transform string to int", err)
	}
	h.Model.DeleteToDo(i)
}

func (h *Handle) UpdateToDo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	todo := model.Todo{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		fmt.Println("ERROR decoding JSON - ", err)
		return
	}
	i, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Println("error bad transform string to int", err)
	}
	h.Model.UpdateToDo(i, todo.Name, todo.Description)
	defer r.Body.Close()

	//	w.Write([]byte(ps.ByName("id")))&todo.Name,
}

func jsonResponse(w http.ResponseWriter) {
	w.Write([]byte("Success"))
	w.WriteHeader(200)
}

func (h *Handle) FetchToDo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	i, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Println("error bad transform string to int", err)
	}
	fmt.Println(h.Model.FetchToDo(i))

}

func (h *Handle) FetchToDos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(h.Model.FetchToDos())
}

func (h *Handle) CreateToDo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	todo := model.Todo{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		fmt.Println("ERROR decoding JSON - ", err)
		return
	}
	defer r.Body.Close()
	h.Model.CreateToDo(todo.Name, todo.Description)
	jsonResponse(w)
}

func main() {

	db, err := sql.Open("postgres", "postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable")
	if err != nil {
		panic(err)
	}

	todosModel := model.TodosModel{
		DB: db,
	}

	router := httprouter.New()
	handle := Handle{
		Model: todosModel,
	}

	router.POST("/api/todos/", handle.CreateToDo)
	router.GET("/api/todos/", handle.FetchToDos)
	router.GET("/api/todos/:id/", handle.FetchToDo)
	router.PATCH("/api/todos/:id/", handle.UpdateToDo)
	router.DELETE("/api/todos/:id/", handle.DeleteToDo)
	log.Fatal(http.ListenAndServe(":8080", router))
}

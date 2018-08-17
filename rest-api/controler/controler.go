package controler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/islanzari/go-exercises/rest-api/model"
	"github.com/julienschmidt/httprouter"
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

	h.Model.CreateToDo(todo.Name, todo.Description)
	jsonResponse(w)
}

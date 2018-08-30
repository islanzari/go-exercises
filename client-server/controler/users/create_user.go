package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/islanzari/go-exercises/client-server/request"
)

func (h Handle) CreateUser(w http.ResponseWriter, r *http.Request) {

	var data struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
		Email   string `json:"email"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println("Error while decoding request bodu", err)
		request.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	user, err := h.Users.CreateUser(data.Name, data.Surname, data.Email)
	if err != nil {
		log.Println("Error while creating todo", err)
		request.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	request.Success(w, user)
	log.Println("created user")
}

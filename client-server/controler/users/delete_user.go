package users

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/islanzari/go-exercises/client-server/request"
)

func (h Handle) DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	i, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Error bad convert string to int")
		request.Error(w, "Invalid string", http.StatusBadRequest)
		return
	}

	err = h.Users.DeleteUser(uint64(i))
	if err != nil {
		log.Println("Error while creating todo")
		request.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	request.Success(w, nil)
	log.Println("delete user")
}

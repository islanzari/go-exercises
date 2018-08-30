package users

import "github.com/islanzari/go-exercises/client-server/model"

type Users interface {
	CreateUser(name, surname, email string) (model.User, error)
	DeleteUser(id uint64) error
	GetUser(id uint64) (model.User, error)
}

type Handle struct {
	Users Users
}

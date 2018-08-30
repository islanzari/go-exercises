package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

// User represents single user
type User struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

// Database is an object containing many users
type Database struct {
	ID    uint64 `json:"lastID"`
	m     sync.RWMutex
	Users map[uint64]User `json:"users"`
}

// ErrRecordNotFound comunicate when records are not found
var ErrRecordNotFound = errors.New("record not found")

// DeleteUser delete user by id
func (d *Database) DeleteUser(id uint64) error {
	d.m.Lock()
	defer d.m.Unlock()
	_, ok := d.Users[id]
	if ok {
		delete(d.Users, id)
	} else {
		return ErrRecordNotFound
	}

	return nil
}

// CreateUser this add user to database and count CounterID
func (d *Database) CreateUser(name string, surname string, email string) (User, error) {
	d.m.Lock()
	d.ID++
	user := User{ID: d.ID, Name: name, Surname: surname, Email: email}
	d.Users[d.ID] = user
	d.m.Unlock()
	return user, nil
}

// GetUser Geting user and show him in console
func (d *Database) GetUser(id uint64) (User, error) {
	d.m.Lock()
	user, ok := d.Users[id]
	if ok {
		d.m.Unlock()
		return user, nil
	}
	d.m.Unlock()
	return User{}, ErrRecordNotFound
}

// New use function timeSave to convert type database to json and save this data in json file
func New(t time.Duration) *Database {
	usersModel := &Database{
		Users: make(map[uint64]User),
	}
	go usersModel.timeSave(t)
	return usersModel
}

func (d *Database) timeSave(t time.Duration) {
	for {
		d.m.Lock()
		jsonString, err := json.Marshal(d)
		if err != nil {
			log.Println(err)
		}
		ioutil.WriteFile("database.json", jsonString, 0644)
		log.Println("zapisano zmiany")
		d.m.Unlock()
		time.Sleep(t)
	}
}

// LoadModel read database from file
func (d *Database) LoadModel(s string) error {
	content, err := ioutil.ReadFile(s)
	if err != nil {
		return err
	}
	d.m.Lock()
	err = json.Unmarshal(content, &d)
	d.m.Unlock()
	return err
}

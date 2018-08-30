package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/islanzari/go-exercises/client-server/cli/client"
)

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)

	for {
		var action string
		_, err := fmt.Scan(&action)
		if err != nil {
			log.Println(err)
		}
		switch action {
		case "adduser":
			err = client.Login()
			if err != nil {
				log.Println(err)
			}
		case "deleteuser":
			err = client.DeleteUser()
			if err != nil {
				log.Println(err)
			}
		case "getuser":
			err = client.GetUser()
			if err != nil {
				log.Println(err)
			}
		default:
			fmt.Println("Unknown task", action)
		}
	}
}

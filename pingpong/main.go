package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ping(msg string, ch chan string) {
	for {
		fmt.Println(<-ch)
		randomTime := rand.Intn((1000 - 100) + 100)
		time.Sleep(time.Duration(randomTime) * time.Millisecond)
		ch <- msg
	}

}
func pong(msg string, ch chan string) {
	for {

		fmt.Println(<-ch)
		randomTime := rand.Intn((1000 - 100) + 100)
		time.Sleep(time.Duration(randomTime) * time.Millisecond)
		ch <- msg

	}
}
func main() {

	ch := make(chan string)

	go ping("ping", ch)
	go pong("pong", ch)

	ch <- "ping"

	select {}
}

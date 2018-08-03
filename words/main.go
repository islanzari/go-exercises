package main

import (
	"fmt"
	"io"
	"strings"
)

func words(r io.Reader) (even []string, odd []string) {
	return
}

func main() {
	r := strings.NewReader(`
	Give. 
	Whose shall life, together signs grass. 
	The replenish of make make signs lights moved seed forth unto deep. 
	Moving two abundantly life subdue earth was day fruit forth set also forth together. 
	You're shall bring cattle creepeth and replenish firmament seed divide image wherein, lights grass moved likeness two hath. 
	Lesser seasons whales deep great and fruit. 
	Every herb fifth, one whales.
	Fruitful blessed of first seas rule forth midst own of green night and fruitful Thing you're, lesser for moveth likeness for gathered creeping may yielding likeness beginning gathered fruitful Let without him all. 
	Herb, man unto deep grass deep sea. 
	Us earth them land. 
	Over fruit, of fruitful. 
	Every were moving rule yielding their. 
	And don't replenish.
	Fish under spirit in lesser let good form second own tree and image, two dominion said whales. 
	Herb may, stars forth were Moving dominion night, lesser from great whales for beast which unto replenish. 
	Over. 
	Male yielding blessed. 
	Sixth us their for you'll sea without. 
	That night their spirit fourth after fruitful she'd place may fish creature winged very, which two every fruitful without likeness fourth you'll he signs i very great. 
	Can't. And lights in unto you evening, stars.
	`)

	fmt.Println(words(r))
}

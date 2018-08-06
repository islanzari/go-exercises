package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func splitStringToFile(fileName string, slice []string) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Println("can't open file")
	}
	for _, v := range slice {
		_, err = f.WriteString(v + " ")

	}
}
func counts(fields []string) (even []string, odd []string) {
	var count int
	for _, v := range fields {
		for _, r := range v {
			switch r {
			case 'a', 'e', 'ą', 'ę', 'o', 'ó', 'y', 'u', 'i', 'A', 'E', 'Ą', 'Ę', 'O', 'Ó', 'Y', 'U', 'I':
				count++
			}
		}
		if count%2 == 0 {
			if count == 0 {
				odd = append(odd, v)
			} else {
				even = append(even, v)
			}
		}
		if count%2 == 1 {
			odd = append(odd, v)
		}
		count = 0
	}
	return even, odd
}

func words(r io.Reader) (even []string, odd []string) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fields := strings.Fields(string(b))
	evenWords, oddWordls := counts(fields)
	splitStringToFile("even.txt", evenWords)
	splitStringToFile("odd.txt", oddWordls)
	return evenWords, oddWordls
}

func main() {
	var file string
	flag.StringVar(&file, "file", "default", "a bool var")
	flag.Parse()

	f, err := os.OpenFile("lorem.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Println("can't open file lorem.txt")
	}

	fmt.Println(words(f))
	if err := f.Close(); err != nil {
		log.Println("can't print function words")
	}
}

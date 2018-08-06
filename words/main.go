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

func splitStringToFile(fileName string, slice []string) error {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	for _, v := range slice {
		_, err = f.WriteString(v + " ")
		if err != nil {
			return err
		}
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
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

func words(r io.Reader) (even []string, odd []string, err error) {

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, nil, err
	}

	fields := strings.Fields(string(b))
	evenWords, oddWordls := counts(fields)
	splitStringToFile("even.txt", evenWords)
	if err != nil {
		return nil, nil, err
	}
	splitStringToFile("odd.txt", oddWordls)
	if err != nil {
		return nil, nil, err
	}
	return evenWords, oddWordls, err
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

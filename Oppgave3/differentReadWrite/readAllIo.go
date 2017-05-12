package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Åpner filen for lesing
	file, err := os.Open("leking.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Leser filen
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", data)
}

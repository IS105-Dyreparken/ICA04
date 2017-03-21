package main

import (
	"io/ioutil"
	"log"
)

// main Dette programmet lager/åpner, skriver og lukker til en fil.
// Filen som blir brukt heter "leking.txt"
// Teksten som blir skrevet er "Boop"
func main() {
	err := ioutil.WriteFile("leking.txt", []byte("Boop\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

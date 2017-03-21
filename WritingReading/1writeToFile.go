package main

import (
	"log"
	"os"
)

func main() {
	// Åpner en fil for skriving
	file, err := os.OpenFile(
		"leking.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Skriver bytes til filen
	byteSlice := []byte("Jeg spiste pizza i dag\n")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}

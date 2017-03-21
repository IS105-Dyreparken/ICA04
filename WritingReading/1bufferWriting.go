package main

import (
	"bufio"
	"log"
	"os"
)

// main Funksjonen skriver til en fil ved bruk av buffer
func main() {
	// Åpner filen får lesing
	file, err := os.OpenFile("leking.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Lager en buffer skriver for filen
	bufferedWriter := bufio.NewWriter(file)

	// Skriver tekst til bufferen
	bytesWritten, err := bufferedWriter.WriteString(
		"Benjamin\n" +
			"Sindre\n" +
			"Shiwwan\n" +
			"Nicolai\n",
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)

	bytesWritten, err = bufferedWriter.WriteString(
		"Marius\n" +
			"Ella\n" +
			"Erik\n",
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)
	// Skriver fra memory til disk
	bufferedWriter.Flush()
}

package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"ICA04/Oppgave1/lineshift"
)

func main() {
	fileToByteslice("text1.txt")
	fileToByteslice("text2.txt")
}

func fileToByteslice(filename string) []byte {

	// Open file for reading
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	sizeOfSlice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, sizeOfSlice)

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%c", byteSlice)
	fmt.Printf("%x", byteSlice)
	lineshift.Lineshift(byteSlice)
	return (byteSlice)
}

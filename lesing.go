package main

import (
	"io"
	"log"
	"os"
  "fmt"

)



func main() {
  filepath := "/Users/sindrethompson/Desktop/GO/src/ICA04/is105-ica04/files/text1.txt"
  filepath2 := "/Users/sindrethompson/Desktop/GO/src/ICA04/is105-ica04/files/text2.txt"
	fileToByteslice(filepath)
  fileToByteslice(filepath2)
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
	size_of_slice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, size_of_slice)

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

  //fmt.Println(byteSlice)
    fmt.Printf("%c", byteSlice)
    //fmt.Printf("%s", byteSlice)
    fmt.Printf("% X\n", byteSlice)
  return (byteSlice)
}

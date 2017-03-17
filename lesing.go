package main

import (
	"io"
	"log"
	"os"
  	"fmt"
	"strings"
)



func main() {
  	filepath := "./files/text1.txt"
  	filepath2 := "./files/text2.txt"
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
	//fmt.Printf("%c", byteSlice)
    	//fmt.Printf("%s", byteSlice)
    	//fmt.Printf("% X\n", byteSlice)
		Lineshift(byteSlice)
  	return (byteSlice)
}


func Lineshift(inputArray []byte) {

	//Finner lengden på arrayet og justerer indeksene slik at man får siste og nest siste posisjon i byte-arrayet
	lineFirst := len(inputArray)
	lineLast := len(inputArray)
	lineLast --
	lineFirst --
	lineLast --

	//Henter ut bytes på siste og nest siste plass i arrayet og printer disse + hele arrayet, mest for debugging
	arrFirst := inputArray[lineFirst]
	arrLast := inputArray[lineLast]
	//fmt.Printf("% X\n", arrFirst)
	//fmt.Printf("% X\n", arrLast)
	//fmt.Printf("% X\n", inputArray)
	//fmt.Println() //Printer en tom linje for å lage litt plass

	//Sammenligner de uthentede bytene fra arrayet med kode for CRLF og LF linjeskift
	x := fmt.Sprintf("% X\n", arrFirst)
	z := fmt.Sprintf("% X\n", arrLast)

	var a = "A"
	var b = "D"

	//Fjerner først "whitespace" (newline, space osv), og sjekker deretter om strengen er stemmer overens med forskjellige formater for linjeskift
	if strings.TrimSpace(x) == a && strings.TrimSpace(z) == b {
		fmt.Println("Format = CRLF")
	} else if strings.TrimSpace(x) == a && strings.TrimSpace(z) != b {
		fmt.Println("Format = LF or other similar")
	} else {
		fmt.Println("Format = Unknown")
	}

}
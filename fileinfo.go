package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo os.FileInfo
	err      error
)

func main() {
	//Stat returns file info. It will return
	// an error if there is no file.
	fmt.Print("Skrive inn file navn: ")
	var input string
	fmt.Scanln(&input)

	fileInfo, err = os.Stat(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())  // størrelse på x in byest, kibibytes, mibibytes og gibibytes
	fmt.Println("Is Directory: ", fileInfo.IsDir()) // hvis det er en mappe bli det true
	if fileInfo.Mode().IsRegular() {                //
		fmt.Println("Is a regular: true") // hvis det er en file, bli det true
	} else {
		fmt.Println("Is a regular: false") // hvis det er bare en mappe bli det false
	}
	fmt.Println("Permissions:", fileInfo.Mode())

	if fileInfo.Mode()&os.ModeAppend == os.ModeAppend { // Deler filen data/info med andre filer
		fmt.Println("Is Append: true")
	} else {
		fmt.Println("Is Append: false")
	}
	if fileInfo.Mode()&os.ModeDevice == os.ModeDevice { // En fil som lar software kommunisere med
		fmt.Println("Is a device file: true") // "device driver" ved hjelp av standar
	} else { // "input/output system calls"
		fmt.Println("Is a device file: false")
	}
	if fileInfo.Mode()&os.ModeCharDevice == os.ModeCharDevice { //Dette må sjekke om det er riktig ??
		fmt.Println("Is a Unix character: true")
	} else {
		fmt.Println("Is Not Unix character: false")
	}
	fmt.Println("Last modified:", fileInfo.ModeTime())

	if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink { // Det er en symlink
		fmt.Println("Is a Symlink: true")
	} else {
		fmt.Println("Is Not a Symlink: false")
	}
}

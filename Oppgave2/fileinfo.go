package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filnavn := os.Args[1]
	fileInfo(filnavn)
}

func fileInfo(filnavn string) { //tar inn string inn og finner info om filen

	fileInfo, err := os.Lstat(filnavn)
	if err != nil {
		log.Fatal(err)
	}

	//Størrelse i forskjellig former:
	i64 := fileInfo.Size()
	bytes := float64(i64)
	kibi := bytes / 1024
	mibi := kibi / 1024
	gibi := mibi / 1024

	//Printer ut informasjon om filen
	fmt.Println("Infomation about a file:", filnavn)
	fmt.Println("Bytes: ", bytes)
	fmt.Println("Kibibytes: ", kibi)
	fmt.Println("Mibibytes: ", mibi)
	fmt.Println("Gibibytes: ", gibi)

	if fileInfo.Mode().IsDir() == true { // hvis det er en mappe bli det true
		fmt.Println("Is a directory")
	} else if fileInfo.Mode().IsDir() == false {
		fmt.Println("Is not a directory")

		if fileInfo.Mode().IsRegular() { //
			fmt.Println("Is a regular: true") // hvis det er en file, bli det true
		} else {
			fmt.Println("Is a regular: false") // hvis det er bare en mappe bli det false
		}
		fmt.Println("Has Unix permissions:", fileInfo.Mode().Perm()) // Hvilke retigheter man har på filen
		if fileInfo.Mode()&os.ModeAppend == os.ModeAppend {          // Deler filen data/info med andre filer
			fmt.Println("Is append only: true")
		} else {
			fmt.Println("Is not append only: false")
		}
		if fileInfo.Mode()&os.ModeDevice == os.ModeDevice { // En fil som lar software kommunisere med "device driver" ved hjelp av standar input/output system calls
			fmt.Println("Is a device file: true")
		} else {
			fmt.Println("Is a device file: false")
		}
		if fileInfo.Mode()&os.ModeCharDevice == os.ModeCharDevice {
			fmt.Println("Is a Unix character: true")
		} else {
			fmt.Println("Is not Unix character: false")

			if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink { // Det er en symlink
				fmt.Println("Is a Symlink: true")
			} else {
				fmt.Println("Is not a Symlink: false")
			}
		}
	}
}

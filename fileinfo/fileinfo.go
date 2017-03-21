package main

import (
	"fmt"
	"os"
)


func main() {
	fmt.Println("Enter the path to the file: ")
	//var input string
	var input = "./files/text1.txt"
	//fmt.Scanln(&input)

	file, _ := os.Stat(input)

	fmt.Println("File name:", file.Name())      // Returnerer filenavnet på valgt fil
	fmt.Println("Size in bytes:", file.Size())  //størrelse
	fmt.Println("Is Directory: ", file.IsDir()) // Er det en mappe?

	regornot := file.Mode().IsRegular()
		if regornot == true {
			fmt.Println("This file is a regular file")
		} else if regornot == false {
			fmt.Println("This file is not a regular file")
		}
	/**
	Hva skjer her??
	if fileInfo.Mode()&os.ModeAppend == os.ModeAppend { // Deler filen data/info med andre filer
		fmt.Println("Is Append: true")
	} else {
		fmt.Println("Is Append: false")
	}
	**/

	if file.Mode.()&os.ModeDevice == os.ModeDevice { // En fil som lar software kommunisere med "device driver" vha standard input/output systemkall"
		fmt.Println("Is a device file: true")
	} else {
		fmt.Println("Is a device file: false")
	}

	fmt.Println("Has unix permissions: ", file.Mode())      // Hvilke rettigheter man har på filen
	fmt.Println("Last modified: ", file.ModTime()) // Tiden hvor det sist ble gjort endring(er) i denne filen
}

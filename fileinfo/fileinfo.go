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
	fmt.Print("Enter name of the file: ")
	var input string
	fmt.Scanln(&input)

	fileInfo, err = os.Stat(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File name:", fileInfo.Name())      //fil navn
	fmt.Println("Size in bytes:", fileInfo.Size())  //størrelse
	fmt.Println("Is Directory: ", fileInfo.IsDir()) // Er det en mappe
	{
		if fileInfo.Mode().IsRegular() { // vanlig fil eller ikke
			fmt.Println("Is a regular: true")
		} else {
			fmt.Println("Is a regular: false")
		}
	}
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

	fmt.Println("Permissions:", fileInfo.Mode())      // Hvilke rettigheter har man til filen
	fmt.Println("Last modified:", fileInfo.ModTime()) // sist endring gjort på denne filen

}

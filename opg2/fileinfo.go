package main

import (
	"fmt"
	"os"
)

func main() {

	fileInfo, err := os.Lstat(os.Args[1])
	if err != nil {
		fmt.Println("can't read file:", os.Args[1])
		panic(err)
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

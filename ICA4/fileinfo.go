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
	// Stat returns file info. It will return
	// an error if there is no file.
	fileInfo, err = os.Stat("fileinfo.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	{
		if fileInfo.IsDir() {
			fmt.Println("it's a directory")
		} else {
			fmt.Println("it's not a directory")
		}
	}
	if !fileInfo.Mode().IsRegular() {
		fmt.Println("it's regular")
	} else {
		fmt.Println("it's not regular")
	}

	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())

	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())

}

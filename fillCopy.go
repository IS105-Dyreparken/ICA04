package main

import (
    "log"
    "os"
    "bufio"
    "fmt"
)

var (
    newFile *os.File
    err     error
)

func main() {
  //createFile()
  readText1()
}

func createFile() {
    newFile, err = os.Create("nyFil.txt")
    if err != nil {
        log.Fatal(err)
    }
    log.Println(newFile)
    newFile.Close()
}

func readText1() {
  file, err := os.Open("text1.txt")
  if err != nil {
      log.Fatal(err)
  }

  bufferedReader := bufio.NewReader(file)
  dataString, err := bufferedReader.ReadString('\n')
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Read string: %s\n", dataString)

  //newFile, err := os.OpenFile("nyFil.txt", os.O_WRONLY, 0666)
  //if err != nil {
    //log.Fatal(err)
  //}
  //bufferedWriter := bufio.NewWriter(newFile)
  //writeText, err := bufferedWriter.Write(dataString)
}

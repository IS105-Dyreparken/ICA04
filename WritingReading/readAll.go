package main

import (
    "os"
    "log"
    "fmt"
    "io/ioutil"
)

func main() {
    // Open file for reading
    file, err := os.Open("leking.txt")
    if err != nil {
        log.Fatal(err)
    }

    // os.File.Read(), io.ReadFull(), and
    // io.ReadAtLeast() all work with a fixed
    // byte slice that you make before you read

    // ioutil.ReadAll() will read every byte
    // from the reader (in this case a file),
    // and return a slice of unknown slice
    data, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s\n", data)
}

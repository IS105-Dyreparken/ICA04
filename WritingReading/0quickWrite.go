package main

import (
    "io/ioutil"
    "log"
)

func main() {
    err := ioutil.WriteFile("leking.txt", []byte("Boop\n"), 0666)
    if err != nil {
        log.Fatal(err)
    }
}

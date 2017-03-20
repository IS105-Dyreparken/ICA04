package main

import (
    "bufio"
    "fmt"
    "os"
)

func lineCount(filename string) (int64, error) {
    lc := int64(0)
    f, err := os.Open("leking.txt")
    if err != nil {
        return 0, err
    }
    defer f.Close()
    s := bufio.NewScanner(f)
    for s.Scan() {
        lc++
    }
    return lc, s.Err()
}

func main() {
    filename := `testfile`
    lc, err := lineCount(filename)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(filename+" line count:", lc)
}

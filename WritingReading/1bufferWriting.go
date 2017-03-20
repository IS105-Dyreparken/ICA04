package main

import (
    "log"
    "os"
    "bufio"
)

func main() {
    // Open file for writing
    file, err := os.OpenFile("leking.txt", os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Create a buffered writer from the file
    bufferedWriter := bufio.NewWriter(file)

    // Write string to buffer
    bytesWritten, err := bufferedWriter.WriteString(
        "Benjamin\n" +
        "Sindre\n" +
        "Shiwwan\n" +
        "Nicolai\n",
    )
    if err != nil {
        log.Fatal(err)
}
    log.Printf("Bytes written: %d\n", bytesWritten)

    // Write memory buffer to disk
    //bufferedWriter.Flush()

    bytesWritten, err = bufferedWriter.WriteString(
        "Marius\n" +
        "Ella\n" +
        "Erik\n",
    )
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Bytes written: %d\n", bytesWritten)

    bufferedWriter.Flush()
}

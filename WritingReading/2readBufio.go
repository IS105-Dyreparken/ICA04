package main

import (
	"bufio"
	"fmt"
	"os"
)

// main Funksjon som skal finne antall linjer lest
// og fem mest leste runer.
func main() {
	lineCount()
}

func lineCount() {
	file, _ := os.Open("leking.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	fmt.Println("Number of lines:", lineCount)
}

func findRunes() {

}

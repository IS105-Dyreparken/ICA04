package main

import (
	"bufio"
	"fmt"
	"os"
)

// main Funksjon skal finne antall linjer
func main() {
	file, _ := os.Open("pg100.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	fmt.Println("Number of lines:", lineCount)
}

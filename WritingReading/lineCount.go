package main

import (
	"bufio"
	"fmt"
	"os"
)

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

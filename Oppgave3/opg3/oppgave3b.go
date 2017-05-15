// Dette er ikke vår kode, kode tatt ifra:
// https://github.com/Zwirc/IS-105/blob/master/ICA04/src/oppgaver/oppgave3.go
// Se readme for forklaring av kode
package opg3

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
)

// Oppgave3b bruker 3 funksjoner for å
// finne linjeskift og runer.
func Oppgave3b() {
	fmt.Println("Skriv inn filen du ønsker å scanne:")
	fmt.Println("Tekstfil i mappe: pg100.txt")
	var filelenght = bufio.NewScanner(os.Stdin)
	var fileinput = ""
	for filelenght.Scan() {
		fileinput = string(filelenght.Text())
		break
	}

	fmt.Print("Antall linjeskift i filen: ")
	fmt.Println(findInText(fileinput, "\n"))

	// Henter map av filen
	m := countFile(fileinput)
	// Sorterer og printer resultat:
	sortAndPrint(m)

}

// Oppgave 3b (1/3)
// Finner tekst i fil.
func findInText(filename string, find string) int {
	// Last inn filen
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Stor buffer for å håndtere store filer
	buf := make([]byte, 32*1024)
	count := 0
	// Gjør om string til byte for søk
	search := []byte(find)
	// Leter etter tekst:
	c, _ := file.Read(buf)
	count += bytes.Count(buf[:c], search)
	return count
}

// Oppgave 3b (2/3)
// Lager en map med resultatet for hver rune, og returnerer denne
func countFile(filename string) map[int]string {
	m := make(map[int]string)
	// Runer for Ascii code valgt:
	for i := 0x20; i <= 0x7F; i++ {
		count := findInText(filename, string(i))
		rune := string(i)
		m[count] = rune
	}
	fmt.Println()
	return m
}

// Oppgave 3b (3/3)
// Sorterer map listen med append, og printer ut de fem største verdier.
func sortAndPrint(m map[int]string) {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// Print 5 største:
	fmt.Println("5 mest brukte runes:")
	value1 := keys[len(keys)-1]
	value2 := keys[len(keys)-2]
	value3 := keys[len(keys)-3]
	value4 := keys[len(keys)-4]
	value5 := keys[len(keys)-5]
	fmt.Println("Antall:", value1, "Rune:", m[value1])
	fmt.Println("Antall:", value2, "Rune:", m[value2])
	fmt.Println("Antall:", value3, "Rune:", m[value3])
	fmt.Println("Antall:", value4, "Rune:", m[value4])
	fmt.Println("Antall:", value5, "Rune:", m[value5])
}

//func Oppgave3c(int int) {
//fmt.Println("Kjør benchmark go test -bench=. i benchmark mappen")
//if int == 1 {
//fileInfo("pg100.txt")
//}
//}

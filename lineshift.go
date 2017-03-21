package lineshift

import (
	"fmt"
	"strings"
)

func Lineshift(inputArray []byte) {

	//Finner lengden på arrayet og justerer indeksene slik at man får siste og nest siste posisjon i byte-arrayet
	lineFirst := len(inputArray)
	lineLast := len(inputArray)
	lineLast--
	lineFirst--
	lineLast--

	//Henter ut bytes på siste og nest siste plass i arrayet og printer disse + hele arrayet, mest for debugging
	arrFirst := inputArray[lineFirst]
	arrLast := inputArray[lineLast]
	//fmt.Printf("% X\n", arrFirst)
	//fmt.Printf("% X\n", arrLast)
	//fmt.Printf("% X\n", inputArray)
	//fmt.Println() //Printer en tom linje for å lage litt plass

	//Sammenligner de uthentede bytene fra arrayet med kode for CRLF og LF linjeskift
	x := fmt.Sprintf("% X\n", arrFirst)
	z := fmt.Sprintf("% X\n", arrLast)

	var a = "A"
	var b = "D"

	//Fjerner først "whitespace" (newline, space osv), og sjekker deretter om strengen er stemmer overens med forskjellige formater for linjeskift
	if strings.TrimSpace(x) == a && strings.TrimSpace(z) == b {
		fmt.Println("Format = CRLF")
	} else if strings.TrimSpace(x) == a && strings.TrimSpace(z) != b {
		fmt.Println("Format = LF or other similar")
	} else {
		fmt.Println("Format = Unknown")
	}

}

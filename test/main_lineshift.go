package main

import (
  "fmt"
)




func main(){
  Lineshift()
}
const two = 2
func Lineshift() {
  testStr := "Jesus lever\n"
  byteArray := []byte(testStr)

  lineFirst := len(byteArray)
  lineLast := len(byteArray)
  lineLast - two	//Trekker fra to slik at vi får den nest siste karakteren i arrayet
  lineLast --	//Trekker fra en slik at vi får den siste karakteren i arrayet

  //fmt.Println(lineFirst) 	//Test for å se om lineFirst returnerer riktig verdi
  //fmt.Println(lineLast)	//Samme som over, bare for lineLast

  fmt.Printf("%X\n", byteArray[lineFirst])
  fmt.Printf("%X\n", byteArray[lineLast])
  fmt.Printf("% X\n", byteArray)
  //return byteArray(lineIndex)
}

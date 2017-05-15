// Dette er ikke vår kode, kode tatt ifra:
// https://github.com/Zwirc/IS-105/blob/master/ICA04/src/oppgaver/oppgave3.go
// Se readme for forklaring av kode
package main

import "testing"

func Benchmark3c(b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StartTimer()
		opg3.Oppgave3c(1)
		b.ReportAllocs()
		b.StopTimer()

	}

}

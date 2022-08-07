package logic01

import (
	"fmt"
	"testing"
)

func TestLogic01Soal1000(t *testing.T) {
	n := 12
	a := 1 // bilangan yg akan dikalikan
	b := 4 // pengali bilangan
	for i := 1; i <= n; i++ {
		if i%4 == 0 { // kondisi tiap index ke-4, untuk merubah kondisi (a) ke tahap selanjutnya, dan (b) ke kondisi awal
			fmt.Print(999, "\t")
			a += 1 // a dinaikan +1 agar bilangan jadi 2
			b = 4  // b direset ke 4 untuk dikurangi lagi, agar pengali jadi 3,2,1
		} else {
			b -= 1
			fmt.Print(a*b, "\t")
		}
	}
}

package logic01

import (
	"fmt"
	"math"
	"testing"
)

/*
Soal 06
nilai n	15
index	1	2	3	4	5	6	7	8	9	10	11	12	13	14	15
value	3	6	9	16	15	18	21	64	21	18	15	144	9	6	3

5W + 1H method :
what? print given sequence of number (1 dimmensional array with n:= 15)
who? 3 starting variable : n = 15, nilai tengah = n/2, a = 3
when? 2 pattern : when i %4 == 0 print i^2, a+=3 when i <= nilai tengha & a-=3 when i >= nilai tengah
where? conditoning  (i%4 == 0) inside the loop multiple of 4 (4,8,12)
why? to create 1 dimensional array using 1 for loop + if method
how? using for loop method with 2 condition (if & else) to create 2 pattern in 1 dimensional array

computational thinking method :
Main problem :
print 1 dimensionall array n = 15, created using for loop
2 starting number : 1. a = 0, 2. middle value = n / 2
decompose :
a = 0+3
*/

func TestLogic01Soal0600(t *testing.T) {
	n := 15
	nt := n / 2
	angka := 3

	// loop baris:start
	for i := 1; i <= n; i++ {
		if i%4 == 0 { // ketika i habis dibagi 4, maka :
			fmt.Print(math.Pow(float64(i), 2), "\t") // print pangkat 2 dari i, lalu tab
		} else { // ketika i tidak habis dibagi 4, maka :
			fmt.Print(angka, "\t") // print angka saja
		}
		if i <= nt { // pengujian i harus independen, agar angka a bisa terus ditambah/dikurang tanpa tergantung kondisi modulus dari i
			angka += 3 // tambah nilai angka ketika i <= nilai tengah
		} else {
			angka -= 3 // kurangi nilai angka ketika i >= nilai tengah
		}
	}
}

func TestLogic01Soal0601(t *testing.T) {
	n := 15
	mid := n / 2
	a := 3
	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			fmt.Print(math.Pow(float64(i), 2), "\t")
		} else {
			fmt.Print(a, "\t")
		}
		if i <= mid {
			a += 3
		} else {
			a -= 3
		}
	}
}

func TestLogic01Soal0602(t *testing.T) {
	n := 15
	tengah := n / 2
	awal := 3
	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			fmt.Print(math.Pow(float64(i), 2), "\t")
		} else {
			fmt.Print(awal, "\t")
		}
		if i <= tengah {
			awal += 3
		} else {
			awal -= 3
		}
	}
}

func TestLogic01Soal0603(t *testing.T) {
	n := 15
	te := n / 2
	a := 3
	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			fmt.Print(math.Pow(float64(i), 2), "\t")
		} else {
			fmt.Print(a, "\t")
		}
		if i <= te {
			a += 3
		} else {
			a -= 3
		}
	}
}
func TestLogic01Soal0604(t *testing.T) {
	n := 15
	mi := n / 2
	aw := 3
	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			fmt.Print(math.Pow(float64(i), 2), "\t")
		} else {
			fmt.Print(aw, "\t")
		}
		if i <= mi {
			aw += 3
		} else {
			aw -= 3
		}
	}
}

func TestLogic01Soal0605(t *testing.T) {
	n := 15
	a := n / 2
	b := 3
	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			fmt.Print(math.Pow(float64(i), 2), "\t")
		} else {
			fmt.Print(b, "\t")
		}
		if i <= a {
			b += 3
		} else {
			b -= 3
		}
	}
}

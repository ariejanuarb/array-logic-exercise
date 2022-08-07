package logic01

import (
	"fmt"
	"math"
	"testing"
)

/*
Soal_09
nilai n	12
index	1	2	3	4	5	6	7	8	9	10	11	12
value	1	2	3	1	4	9	1	8	27	1	16	81
		1^1 2^1 3^1|1^2 2^2 3^2|1^3 2^3 3^3|1^4 2^4 3^4
*/

func TestLogic01Soal0900(t *testing.T) {
	n := 12 // jumlah n pada deret
	a := 1  // bilangan pengali
	b := 1  // bilangan pangkat
	for i := 1; i <= n; i++ {
		fmt.Print(math.Pow(float64(a), float64(b)), "\t") // urutan hasil = 1 dari 1^1, 2 dari 2^1, 3 dari 3^1
		if i%3 == 0 {                                     // saat i kelipatan 3, maka : kembalikan a ke =1, dan tambah b+=1
			a = 1
			b += 1 // namun ketika i = 3,6,9,12, kondisi b bertambah 1 secara berulang/bertahap
		} else { // saat i di kelipatan yg lainnya, maka : tambah nilai a += 1
			a += 1
		}
	} // ketika loop berakhir, kondisi b pada i = 1,2,3 adalah 1. sementara kondisi a pada i = 1,2,3 adalah 1
}

func TestLogic01Soal0901(t *testing.T) {
	n := 12
	c := 1
	d := 1
	for i := 1; i <= n; i++ {
		fmt.Print(math.Pow(float64(c), float64(d)), "\t")
		if i%3 == 0 {
			c = 1
			d += 1
		} else {
			c += 1
		}
	}
}

func TestLogic01Soal0902(t *testing.T) {
	n := 12
	bil := 1
	pang := 1
	for i := 1; i <= n; i++ {
		fmt.Print(math.Pow(float64(bil), float64(pang)), "\t")
		if i%3 == 0 {
			bil = 1
			pang += 1
		} else {
			bil += 1
		}
	}
}

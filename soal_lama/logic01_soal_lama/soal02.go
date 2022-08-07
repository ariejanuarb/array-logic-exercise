package main

import "fmt"

/**

Soal 2
n = 9
	0	1	2	3	4	 5	  6	   7	8
0 |	1 |	3 |	5 |	7 |	9 |	11 | 13 | 15 | 17 |
	 +2  +2  +2  +2  +2   +2   +2   +2
	angka diatas adalah deret bilangan ganjil
*/
//solusi 1

func main() {
	n := 9
	a := 1
	for i := 1; i <= n; i++ {
		fmt.Print(a, "\t") // print horizontal untuk membuat array 1 dimensi
		b := a + 2         // menambahkan angka 2 untuk deret angka selanjutnya
		a = b              // menggeser posisi a ke kanan
	}
}

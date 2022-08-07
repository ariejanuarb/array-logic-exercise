package main

import "fmt"

/**
Soal 1
n = 9
	0	1	2	3	4	5	6	7	8		Identifikasi masalah :
0	1								9		1. Array 1 dimensi membutuhkan 1 for, array 2 dimensi butuh 2 nested for
1		2						8			2. membuat baris angka 1 s/d 9 dengan for loop, print ke kanan
2			3				7				3.
3				4		6
4					5
5				4		6
6			3				7
7		2						8
8	1								9
*/

func main() {
	// 2. mengulangi proses ke-1 sebanyak 9x dengan for loop-j
	x := 9
	for j := 1; j <= x; j++ {
		// 1. membuat baris angka 1-9 ke-kanan dengan for loop-i
		n := 9
		for i := 1; i <= n; i++ {
			// 4. menampilkan bilangan diagonal kiri atas - kanan bawah saja 1-1, 2-2, ... ,x-n
			if i == j {
				fmt.Print(ganjil)
				// fmt.Print(j, "-", i, "\t")
				// dari hasil analisis baris diatas, ditemukan pola angka yang sama pada tiap baris
				// pola penjumlahan yang sama dari X = 1+9=10, 2+8=10, 3+7=10, 4+6=10, 5+5=10 .... i+j =10
				// 5. mencetak diagonal dari kanan atas ke kiri bawah dengan kondisi else-if i+j=10
			} else if i+j == 10 {
				fmt.Print(ganjil)
			} else {
				fmt.Print("", "", "", "\t")
			}
			// fmt.Print(i, "\t") // mencetak baris angka
			// fmt.Print(j, "-", i, "\t") // mencetak nomor looping (ke-j) & baris angka 1-9 (ke-i)
		}
		// 3. membuat jeda baris baru ketika loop ke-i selesai
		fmt.Println("\n") // cetak baris baru (kebawah) ketika looping ke-i slesai
	}
}

func ganjil() int {
	n := 9
	a := 1
	for i := 1; i <= n; i++ {
		fmt.Print(a, "\t") // print horizontal untuk membuat array 1 dimensi
		b := a + 2         // menambahkan angka 2 untuk deret angka selanjutnya
		a = b              // menggeser posisi a ke kanan
	}
	return int
}

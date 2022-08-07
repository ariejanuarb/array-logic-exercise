package main

import "fmt"

/**
Soal 3
n = 9
	0	1	2	3	4	5	6	7	8
0	0	2	4	6	8	10	12	14	16
1	0	2						14	16
2	0		4				12		16
3	0			6		10			16
4	0				8				16
5	0			6		10			16
6	0		4				12		16
7	0	2						14	16
8	0	2	4	6	8	10	12	14	16
*/

func main() {
	n := 9

	for j := 1; j <= n; j++ {
		for i := 0; i <= 16; i += 2 {
			//fmt.Print(j, "-", i, "\t")
			if j == 1 || j == 9 {
				fmt.Print(i, "\t")
			} else if i == 0 || i == 16 {
				fmt.Print(i, "\t")
			} else {
				fmt.Print("", "\t")
			}
		}
		fmt.Println("\n") // cetak baris baru (kebawah) ketika looping ke-i slesai
	}
}

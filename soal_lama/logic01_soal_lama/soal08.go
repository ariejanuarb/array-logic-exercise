package main

import "fmt"

/**
Soal 8
n = 9
	0	1	2	3	4	5	6	7	8
0	A	2	C	4	E	6	G	8	I

*/

func main() {
	n := 9
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Print(i, "\t")
		} else {
			switch i {
			case 1:
				fmt.Print("A", "\t")
			case 3:
				fmt.Print("C", "\t")
			case 5:
				fmt.Print("E", "\t")
			case 7:
				fmt.Print("G", "\t")
			case 9:
				fmt.Print("I", "\t")
			}
		}
	}
}

// printing horizontaly to create 1 dimensional array
// jika i genap, print angka genap dari kelipatan 2. 2,4,6,8
// jika i ganjil, print huruf dengan urutan ganjil A C E G I

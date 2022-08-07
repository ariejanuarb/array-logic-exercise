package main

import "fmt"

/**
Soal 7
n = 9
	0	1	2	3	4	5	6	7	8
0	A	B	C	D	E	F	G	H	I
*/

// solusi 1

func main() {
	for n := 'A'; n <= 'I'; n++ {
		fmt.Printf("%c\t", n)
	}
}

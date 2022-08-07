package main

import "fmt"

/**
Soal 9
n = 9
	0	1	2	3		4		5		6		7		8
0	1	3	9	27		81		243		729		2187	6561 	deret bilangan dikali 3 (1 s/d 8)
	1*3=3
*/

func main() {
	n := 9
	a := 1
	for i := 1; i <= n; i++ {
		fmt.Print(a, "\t") // printing horizontaly to create 1 dimensional array
		b := a * 3         // kalikan a dengan 3 dan definisikan b
		a = b              // geser posisi a ke b
	}
}

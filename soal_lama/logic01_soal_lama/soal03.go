package main

import "fmt"

/**

Soal 3
n = 9
	0	1	2	3	4	 5	  6	   7	8
0 |	0 |	2 |	4 |	6 |	8 |	10 | 12 | 14 | 16 |
	  +2  +2  +2  +2  +2   +2   +2   +2
	deret bilangan genap
*/
// solusi 1

func main() {
	n := 9
	a := 0
	for i := 1; i <= n; i++ {
		fmt.Print(a, "\t") // printing horizontaly to create 1 dimensional array
		b := a + 2
		a = b
	}
}

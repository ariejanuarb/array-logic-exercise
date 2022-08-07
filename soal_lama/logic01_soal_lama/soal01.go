package main

import "fmt"

/**

Soal 1
n = 9
	0	1	2	3	4	5	6	7	8
0 |	1 |	2 |	3 |	4 |	5 |	6 |	7 |	8 |	9 |
     +1  +1  +1  +1   +1  +1  +1  +1
*/

func main() {
	n := 10
	for i := 1; i <= n; i++ {
		fmt.Print(i, "\t") // print horizontal untuk membuat array 1 dimensi
	}
}

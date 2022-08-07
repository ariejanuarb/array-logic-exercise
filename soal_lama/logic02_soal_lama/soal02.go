package main

import "fmt"

/**
Soal 2
n = 9
	0	1	2	3	4	5	6	7	8
0	1	3	5	7	9	11	13	15	17 j == 0
1	1								17
2	1								17
3	1								17
4	1								17
5	1								17
6	1								17
7	1								17
8	1	3	5	7	9	11	13	15	17
*/

func main() {
	n := 9
	for j := 0; j < n; j++ {
		a := 1
		for i := 0; i < n; i++ {
			if j == n-n || j == n-1 || i == n-n {
				fmt.Print(a, "\t")
				a += 2
			} else if i < n-1 {
				fmt.Print(" ", "\t")
				a += 2
			} else {
				fmt.Print(a, "\t")
			}
		}
		fmt.Println("\n")
	}
}

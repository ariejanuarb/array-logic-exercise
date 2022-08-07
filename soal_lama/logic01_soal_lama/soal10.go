package main

import "fmt"

/**
Soal 10
n = 9
	0	1	2	3	4	5	6	7	8
0	0	1	8	27	64	125	216	343	512 	bilangan pangkat tiga (0 s/d 8)
*/

func main() {
	n := 9
	for i := 0; i < n; i++ {
		fmt.Print(i*i*i, "\t")
	}
}

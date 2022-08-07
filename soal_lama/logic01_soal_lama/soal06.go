package main

import "fmt"

/**

Soal 1
n = 9
	0	1	2	3	 4	 5	  6	    7	8
0 |	2 |	3 |	5 |	7 |	11 | 13 | 17 | 19 |	23 |
 deret diatas adalah deret bilangan prima
 bilangan prima = bilangan yang > 1 (start from 2)
				dan hanya habis dibagi oleh 1 dan bilangan itu sendiri
				(jika habis dibagi 1 dan bilangan itu sendiri, maka print \t)
*/

// masih bingung, bagaimana caranya supaya bisa pakai n := 9 menggunakan algoritma sieve of Eratosthenes

func main() {
	var num = 29
	for j := 2; j < num; j++ {
		var flag = 0
		for i := 2; i <= j/2; i++ { // ketika j/2 <= i,
			if j%i == 0 { // jika j modulo 0 = 0 adalah true, maka flag =1 dan kembali ke awal loop -> bilangan prima
				flag = 1 // jika false maka bilangan bukan prima, dan lompat line 30
				break
			}
		}
		if flag == 0 { // jika bilangan bukan prima, flag akan ==  1,
			fmt.Print(j, "\t")
		}
	}
}

//func main() {
//	n := 9
//	for i := 0; i < n; i++ {
//
//	}
//}

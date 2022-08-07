package main

import "fmt"

/**

Soal 4
n = 9
		0	1	2	3	4	5	6	 7	  8
  0   |	1 |	1 |	2 |	3 |	5 |	8 | 13 | 21 | 34 |
		1 + 1 = 2 								-> a + b = c, c = a + b
			1 + 2 = 3							-> a = b, b = c (geser posisi)
				2 + 3 = 5
					3 + 5 = 8
						5 + 8 = 13
							8 + 13 = 21
								13 + 21 = 34
*/

// solusi 1
/*func main() {
	n := 9
	a := 1
	b := 1
	c := 0

	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Print(a, "\t")
			continue
		}
		if i == 2 {
			fmt.Print(b, "\t")
			continue
		}
		// swap values
		c = a + b
		a = b // geser posisi a ke b
		b = c // geser posisi b ke c
		fmt.Print(c, "\t")
	}
}*/

//solusi 2
/*func main() {
	n := 9
	a := 1
	b := 1
	for i := 0; i < n; i++ {
		fmt.Print(a, "\t")
		sum := a + b
		a = b // swap
		b = sum
	}
}
*/

func main() {
	n := 12
	a := 3
	b := 1
	for i := 0; i < n; i++ {
		fmt.Print(a, "\t")
		sum := a - b
		a = b // swap
		b = sum
	}
}

package main

import "fmt"

/**

Soal 5
n = 9
		0	1	2	3	4	5	6	 7	  8
  0   |	1 |	1 |	1 |	3 |	5 |	9 | 17 | 31 | 57 |
		1 + 1 + 1 = 3
			1 + 1 + 3 = 5
				1 + 3 + 5 = 9
					3 + 5 + 9 = 17
						5 + 9 + 17 = 31
							9 + 17 + 31 = 57
*/

func main() {
	n := 9
	a := 1
	b := 1
	c := 1
	d := 0

	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Print(a, "\t")
			continue
		}
		if i == 2 {
			fmt.Print(b, "\t")
			continue
		}
		if i == 3 {
			fmt.Print(c, "\t")
			continue
		}
		// swap values
		d = a + b + c
		a = b // geser posisi a ke b
		b = c // geser posisi b ke c
		c = d // geser posisi c ke d
		fmt.Print(c, "\t")
	}
}

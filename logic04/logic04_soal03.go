package logic04

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic04Soal03arr00(n int) {
	// 1 create two dimmensional array 9x9=nxn with dynamic param
	array040201 := array2.NewNumberArray(n, n)
	// 2.1 initiate involved variable and value
	bilangan := 1
	// 2 create nested for loop to iterate array input
	for baris := 0; baris < len(array040201); baris++ {
		for kolom := 0; kolom < len(array040201); kolom++ {
			// 2.2 use for loop variable as array index & insert the value
			// 2.3 write cond. on rows and column to create printing pattern : 4 printing pattern
			// pattern 1 :
			if baris%4 == 0 {
				array040201[n-1-kolom][baris] = int32(bilangan)
				bilangan += 3
			} else if baris%4 == 1 && kolom == n-n {
				array040201[kolom][baris] = int32(bilangan)
				bilangan += 3
			} else if baris%4 == 2 {
				array040201[kolom][baris] = int32(bilangan)
				bilangan += 3
			} else if baris%4 == 3 && kolom == n-1 {
				array040201[kolom][baris] = int32(bilangan)
				bilangan += 3
			}
		}
		// 2. write cond. on rows to create number pattern (1+3, 4+3, n+3)
	}
	// 3 print created array040201
	array2.PrintNumberArrayZeroEmpty(array040201)
}

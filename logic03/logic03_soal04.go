package logic03

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic03Soal04arr00(n int) { // create first 4 line that are dynamic
	//1 create new array
	array0400 := array2.NewNumberArray(n, n)
	//2.1 initialize variable and value : start with 3, +3 in new row. and -3 after middle constraint (n/2)
	number := 3
	middle := n / 2
	//2 create for-loop and use its variable as array index
	for rows := 0; rows < len(array0400); rows++ {
		for column := 0; column < len(array0400); column++ {
			//2.2 insert value to index (val. need parsing)
			if rows%4 == 0 { // rows 0,4,8 so on : start from 0 +4
				array0400[rows][column] = int32(number)
			} else if rows%4 == 1 && column == n-1 { // rows 1,5,9 so on & column n-1 : start from 1 +4
				array0400[rows][column] = int32(number)
			} else if rows%4 == 2 { // rows 2,6,10 so on : start from 2, +4 +4
				array0400[rows][column] = int32(number)
			} else if rows%4 == 3 && column == n-n { // rows 3,7,11 so on & column 0 : start from 3 +4
				array0400[rows][column] = int32(number)
			}
		}
		//2.3 write condition on rows to create number pattern
		if rows < middle {
			number += 3
		} else {
			number -= 3
		}
	}
	//3 print array
	array2.PrintNumberArrayZeroEmpty(array0400)
}

func Logic03Soal04arr01(n int) {
	array0401 := array2.NewNumberArray(n, n)
	num := 3
	mid := n / 2
	for r := 0; r < len(array0401); r++ {
		for c := 0; c < len(array0401); c++ {
			// write condition on rows & column to create 4 Printing pattern
			// pattern 1 : print all number on row 0,4,8,12,n
			if r%4 == 0 { // start from row 0, then added 4, so 0,4,8,12 so-on
				array0401[r][c] = int32(num)
				// pattern 2 : print number only in n-1 column, on row 1,5,9,13,n
			} else if r%4 == 1 && c == n-1 { // start from row 1, added 4, 1,5,9,13 so-on
				array0401[r][c] = int32(num)
				// pattern 3 : print all number on row 2,6,10,14,n
			} else if r%4 == 2 {
				array0401[r][c] = int32(num)
				// pattern 4 : print number only in n-n column, on row 3,7,11,15,n
			} else if r%4 == 3 && c == n-n {
				array0401[r][c] = int32(num)
			}
		}
		// write condition on rows to create number pattern
		// +3 every new row untill mid, after mid, number -3
		if r < mid {
			num += 3
		} else {
			num -= 3
		}
	}
	array2.PrintNumberArrayZeroEmpty(array0401)
}

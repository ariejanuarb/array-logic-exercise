package logic02

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic02Soal07arr00(n int) {
	//1 create two dimmensional array
	array0700 := array2.NewNumberArray(n, n)
	//2 insert value of the array
	//2.1 initialize variable and number used
	a := 3
	//2.2 create for loop to insert value in rows and column
	for baris := 0; baris < len(array0700); baris++ { // rows loop : start
		// a as starting point & reset point on every new row
		for kolom := 0; kolom < len(array0700); kolom++ { // column loop : start
			// 2.3 use baris & kolom variable for array value
			// 2.4 write if condition on rows/column to create desired pattern
			if baris >= kolom { //  baris == kolom || baris > kolom
				array0700[baris][kolom] = int32(a)
			}
		} // column loop : end
		a += 3
	} // rows loop : end
	//3 print array0700
	array2.PrintNumberArrayZeroEmpty(array0700)
}

func Logic02Soal07arr01(n int) {
	//1 create 2 dimmensional array
	array0701 := array2.NewNumberArray(n, n)

	//2 insert array values
	//2.1 initialize variable and values
	number := 3
	//2.2 create for loop rows and column to insert array values
	for rows := 0; rows < len(array0701); rows++ {
		for column := 0; column < len(array0701); column++ {
			//2.3 use rows and column as array index [rows][column] and insert the value
			//2.4 write conditions to create desired pattern
			if rows >= column {
				array0701[rows][column] = int32(number)
			} // no need to write else, because empty number will considered 0, and zero will be printed empty with the printer below
		}
		number += 3 // add 3 right after column loop is finished
	}
	//3 print the array
	array2.PrintNumberArrayZeroEmpty(array0701)
}

func Logic02Soal07arr02(n int) {
	array0702 := array2.NewNumberArray(n, n)
	b := 3
	for r := 0; r < len(array0702); r++ {
		for c := 0; c < len(array0702); c++ {
			if r >= c {
				array0702[r][c] = int32(b)
			}
		}
		b += 3
	}
	array2.PrintNumberArrayZeroEmpty(array0702)
}

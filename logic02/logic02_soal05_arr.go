package logic02

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic02Soal05arr00(n int) {
	// 1. create new empty array
	array0500 := array2.NewNumberArray(n, n)

	//2. insert array value (n,n) using for loop
	value := 3
	nt := n / 2
	for rows := 0; rows < len(array0500); rows++ {
		for column := 0; column < len(array0500); column++ {
			array0500[rows][column] = int32(value)
		}
		// 2.1 condition the row
		if rows < nt {
			value += 3
		} else {
			value -= 3
		}
	}
	// 3. print array0500 filled with value
	array2.PrintNumberArray(array0500)
}

func Logic02Soal05arr01(n int) {
	//1 create array
	array0501 := array2.NewNumberArray(n, n)
	//2 insert number to array
	//2.1 initialize number
	bil := 3
	nilaitengah := n / 2
	//2.2 create loop to insert number
	for baris := 0; baris < n; baris++ {
		for kolom := 0; kolom < n; kolom++ {
			//2.3 insert number to created array
			array0501[baris][kolom] = int32(bil)
		}
		//2.4 write condition to create desired pattern
		if baris < nilaitengah {
			bil += 3
		} else {
			bil -= 3
		}
	}
	//3 print array after finished inserting number
	array2.PrintNumberArray(array0501)
}

func Logic02Soal05arr02(n int) {
	//1 create array
	array0502 := array2.NewNumberArray(n, n)
	//2 insert number to array
	//2.1 intialize number
	a := 3
	mid := n / 2
	//2.2 create loop to insert number in every (n,n)
	for rows := 0; rows < n; rows++ {
		for col := 0; col < n; col++ {
			// 2.3 use loop variable as array element [n][n]
			array0502[rows][col] = int32(a)
		}
		// 2.4 write condition on array rows to create pattern
		if rows < mid {
			a += 3
		} else {
			a -= 3
		}
	}
	//3 print created array0502
	array2.PrintNumberArray(array0502)
}

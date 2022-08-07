package logic02

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic02Soal04arr00(n int) {
	//1 create new array
	array0400 := array2.NewNumberArray(n, n)
	//2 fill array row and col with value
	value := n * 3
	for row := 0; row < len(array0400); row++ {
		for col := 0; col < len(array0400); col++ {
			array0400[row][col] = int32(value) // insert value
		}
		value -= 3 // value -3 every new row
	}
	//3 print array
	array2.PrintNumberArray(array0400)
}

func Logic02Soal04arr01(n int) {
	//1 create new array
	array0401 := array2.NewNumberArray(n, n)
	//2 fill array row and col (n,n) with value using for loop
	value := n * 3
	for rows := 0; rows < len(array0401); rows++ {
		for colu := 0; colu < len(array0401); colu++ {
			// insert value
			array0401[rows][colu] = int32(value)
		}
		value -= 3
	}
	// 3 print array
	array2.PrintNumberArray(array0401)
}

func Logic02Soal04arr02(n int) {
	array0402 := array2.NewNumberArray(n, n)
	bilangan := n * 3
	for baris := 0; baris < len(array0402); baris++ {
		for kolom := 0; kolom < len(array0402); kolom++ {
			array0402[baris][kolom] = int32(bilangan)
		}
		bilangan -= 3
	}
	array2.PrintNumberArray(array0402)
}

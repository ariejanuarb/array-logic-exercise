package logic02

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic02Soal10arr00(n int) {
	//1 create array
	array1000 := array2.NewNumberArray(n, n)
	//3 insert value using for loop
	for r := 0; r < len(array1000); r++ {
		//2 initialize var and val
		num := 3
		for c := 0; c < len(array1000); c++ {
			if r+c <= n-1 {
				array1000[r][c] = int32(num)
			}
			num += 3
		}
	}
	array2.PrintNumberArrayZeroEmpty(array1000)
}

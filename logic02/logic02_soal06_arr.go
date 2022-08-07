package logic02

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic02Soal06arr00(n int) {
	//1 create two dimmensional array
	array0600 := array2.NewNumberArray(n, n)
	// 2 insert value of the array (n,n)
	// 2.1 initialize variable and values
	mid := n / 2
	//2.2 create for loop to insert value in every array elements
	for r := 0; r < len(array0600); r++ { // loop row : start
		num := 3                              // initialize number value & declare every new r loop
		for c := 0; c < len(array0600); c++ { // loop column : start
			// 2.3 use for loop variable as array elements [n][n]
			array0600[r][c] = int32(num) //parse the number from int to int32
			if c < mid {
				num += 3
			} else {
				num -= 3
			}
		} // loop column : end
		//2.4 write condition to create desired pattern
	} // loop row : end
	//3 print the array0600
	array2.PrintNumberArray(array0600)
}


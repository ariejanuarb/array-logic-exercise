package logic04

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic04Soal04arr00(n int) {
	//1. create new array
	array040100 := array2.NewNumberArray(n, n)
	//2.1 initialize variable and constraint value
	num := 1
	//2. create and use for loop row & column variable as array index
	for rows := 0; rows < len(array040100); rows++ {
		for column := 0; column < len(array040100); column++ {
			// 2.2 insert value on array index
			// 2.3 write condition on rows and column to create 4 printing pattern
			if rows%4 == 0 {
				array040100[rows][n-1-column] = int32(num)
				num += 3
			} else if rows%4 == 1 && column == n-n {
				array040100[rows][column] = int32(num)
				num += 3
			} else if rows%4 == 2 {
				// the starting point on row 2,4,6,7 are different from default : it require column conditions to create another printing pattern
				// array[1][2] :: 1 = by default refers to array rows, 2 = by default refers to array column
				// write formula on column index to reverse the printing order : from 0,1,2,3,n to n-1,n-1-1,n-1-1-1 so on
				array040100[rows][column] = int32(num) // if you only write n-1, only 1 column (n-1 = 8) will be affected
				num += 3                               // but if you write [n-1-column] another column will be affected, because "column" variable used as loop variable
			} else if rows%4 == 3 && column == n-1 {
				array040100[rows][column] = int32(num)
				num += 3
			}
		}
	}
	//3. print array
	//fmt.Println(array)
	array2.PrintNumberArrayZeroEmpty(array040100)
}

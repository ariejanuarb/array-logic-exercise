package logic03

import (
	"fmt"
	array2 "github.com/ariejanuarb/golang-logic-dasar/array"
)

func Logic03Soal02(n int) { // sgitiga atas = (logic 08 + logic 10) + logic 05
	a := 3
	nt := n / 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i <= j && i+j <= n-1 {
				fmt.Print(a, "\t")
			} else if i >= j && i+j >= n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Println("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	}
}

func Logic03Soal02arr00Step01(n int) { // create top triangle
	//1 create array
	arr0200 := array2.NewNumberArray(n, n)
	//2.1 initialize number and constraint
	num := 3
	constraint := n / 2
	//2 insert value to array
	for r := 0; r < len(arr0200); r++ {
		for c := 0; c < len(arr0200); c++ {
			// 2.2 write condition on r & c to create printing pattern
			if r <= c && r+c <= n-1 {
				arr0200[r][c] = int32(num)
			}
		}
		// 2.3 write condition on row (outside the loop) to create number pattern
		if r < constraint {
			num += 3
		} else {
			num -= 3
		}
	}
	array2.PrintNumberArrayZeroEmpty(arr0200)
}

func Logic03Soal02arr00Step02(n int) { // create bottom triangle
	//1 create array
	arr0200 := array2.NewNumberArray(n, n)
	//2.1 initialize number and constraint
	num := 3
	constraint := n / 2
	//2 insert value to array
	for r := 0; r < len(arr0200); r++ {
		for c := 0; c < len(arr0200); c++ {
			// 2.2 write condition on r & c to create printing pattern
			if r <= c && r+c <= n-1 {
				arr0200[r][c] = int32(num)
				// 2.2 write condition to create bottom triangle from top triangle
			} else if r >= c && r+c >= n-1 {
				arr0200[r][c] = int32(num)
			}
		}
		// 2.3 write condition on row (outside the loop) to create number pattern
		if r < constraint {
			num += 3
		} else {
			num -= 3
		}
	}
	array2.PrintNumberArrayZeroEmpty(arr0200)
}

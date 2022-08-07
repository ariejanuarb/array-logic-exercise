package logic03

import (
	"fmt"
	array2 "github.com/ariejanuarb/golang-logic-dasar/array"
)

func Logic03Soal01(n int) { // segitiga kiri = (logic 07 + logic 10)  + logic 05
	a := 3
	nt := n / 2
	for i := 0; i < n; i++ { // loop baris:start
		for j := 0; j < n; j++ { // loop kolom:start
			if i >= j && i+j <= n-1 { // logic 8 reversed && logic 10
				fmt.Print(a, "\t")
			} else if i <= j && i+j >= n-1 { // logic 8 && logic 10 reversed
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

func Logic03Soal01arr00Step01(n int) { // create left triangle
	// 1 create array
	array0100 := array2.NewNumberArray(n, n)
	// 2.1 initialize variable and number values
	angka := 3
	nt := n / 2
	// 2 create for loop rows and column
	for baris := 0; baris < len(array0100); baris++ {
		for kolom := 0; kolom < len(array0100); kolom++ {
			// 2.2 write condition to create column pattern
			// using && = true + true = true, both of them must be true
			if baris >= kolom && baris+kolom <= n-1 {
				// 2.3 insert value to rows and column array0100
				array0100[baris][kolom] = int32(angka)
			}
		}
		// 2.4 write condition to create row pattern
		if baris < nt {
			angka += 3
		} else {
			angka -= 3
		}
	}
	// 3 print array0100 with zero number as empty
	array2.PrintNumberArrayZeroEmpty(array0100)
}

func Logic03Soal01arr00Step02(n int) { // mirror the left triangle to right
	// 1 create array
	array0100 := array2.NewNumberArray(n, n)
	// 2.1 initialize variable and number values
	angka := 3
	nt := n / 2
	// 2 create for loop rows and column
	for baris := 0; baris < len(array0100); baris++ {
		for kolom := 0; kolom < len(array0100); kolom++ {
			// 2.2 write condition to create column pattern
			if baris >= kolom && baris+kolom <= n-1 {
				// 2.3 insert value to rows and column array0100
				array0100[baris][kolom] = int32(angka)
			} else if baris <= kolom && baris+kolom >= n-1 { // the triangle mirror
				array0100[baris][kolom] = int32(angka)
			}
		}
		// 2.4 write condition to create row pattern
		if baris < nt { // this written below because first angka need to be printed as 3
			angka += 3
		} else {
			angka -= 3 // if this written at top of the row loop, there will be 0 value on 3 number
		}
	}
	// 3 print array0100 with zero number as empty
	array2.PrintNumberArrayZeroEmpty(array0100)
}

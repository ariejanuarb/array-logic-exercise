package logic02

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic02Soal09arr00(n int) {
	array0900 := array2.NewNumberArray(n, n)
	for b := 0; b < len(array0900); b++ {
		bil := 3
		for k := 0; k < len(array0900); k++ {
			if b+k >= n-1 { // or b+k > n-1 || b+k > n-1
				array0900[b][k] = int32(bil)
			}
			bil += 3 // bil is still add +3 even though the condition are false
		}
	}
	array2.PrintNumberArrayZeroEmpty(array0900)
}

func Logic02Soal09arr01(n int) {
	array0901 := array2.NewNumberArray(n, n)
	for r := 0; r < len(array0901); r++ {
		num := 3 // written here as declaration & reset point values
		for c := 0; c < len(array0901); c++ {
			if r+c >= n-1 {
				array0901[r][c] = int32(num)
			}
			num += 3
		}
	}
	array2.PrintNumberArrayZeroEmpty(array0901)
}

func Logic02Soal09arr02(n int) {
	array0902 := array2.NewNumberArray(n, n)
	for bar := 0; bar < len(array0902); bar++ {
		val := 3
		for kol := 0; kol < len(array0902); kol++ {
			if bar+kol >= n-1 {
				array0902[bar][kol] = int32(val)
			}
			val += 3
		}
	}
	array2.PrintNumberArrayZeroEmpty(array0902)
}

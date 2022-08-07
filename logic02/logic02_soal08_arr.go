package logic02

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic02Soal08arr00(n int) {
	array0800 := array2.NewNumberArray(n, n)
	angka := 3
	for baris := 0; baris < len(array0800); baris++ {
		for kolom := 0; kolom < len(array0800); kolom++ {
			if baris <= kolom {
				array0800[baris][kolom] = int32(angka)
			}
		}
		angka += 3
	}
	array2.PrintNumberArrayZeroEmpty(array0800)
}

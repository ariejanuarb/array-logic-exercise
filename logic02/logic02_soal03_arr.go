package logic02

import (
	array2 "github.com/ariejanuarb/golang-logic-dasar/array"
)

func Logic02Soal03arr00(n int) {
	//1 create array
	arraybaru := array2.NewNumberArray(n, n)
	//2 isi array

	for baris := 0; baris < len(arraybaru); baris++ {
		bil := n * 3 // 2.1 bilangan awal = n * 3 tiap awal baris & tiap baris baru
		for kolom := 0; kolom < len(arraybaru); kolom++ {
			// 2.4 isi array dengan bil
			arraybaru[baris][kolom] = int32(bil)
			bil -= 3
		}
	}
	// 3 print array yg telah diisi
	array2.PrintNumberArray(arraybaru)
}

func Logic02Soal03arr01(n int) {
	//1 create array
	newarray := array2.NewNumberArray(n, n)
	//2 isi array
	for baris := 0; baris < len(newarray); baris++ {
		angka := n * 3
		for kolom := 0; kolom < len(newarray); kolom++ {
			newarray[baris][kolom] = int32(angka)
			angka -= 3
		}
	}
	// 3 print array
	array2.PrintNumberArray(newarray)
}

func Logic02Soal03arr02(n int) {
	arraybaru := array2.NewNumberArray(n, n)
	for rows := 0; rows < len(arraybaru); rows++ {
		value := n * 3
		for colu := 0; colu < len(arraybaru); colu++ {
			arraybaru[rows][colu] = int32(value)
			value -= 3
		}
	}
	array2.PrintNumberArray(arraybaru)
}

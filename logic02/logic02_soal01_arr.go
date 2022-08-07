package logic02

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic02Soal01arr00(n int) {
	// 1. create array
	array := array2.NewNumberArray(n, n)

	// 2. isi array
	angka := 3                        // 2.1 angka pertama = 3
	for i := 0; i < len(array); i++ { // loop baris : start
		for j := 0; j < len(array[i]); j++ { // loop kolom : start
			// isi array
			array[i][j] = int32(angka) // insert angka = 3 pada i = 0++ & j = 0++
		} // loop kolom : end
		angka += 3 // 2.1.1 angka +3 setiap baris baru
	}
	// 3. print array
	array2.PrintNumberArray(array)
}

func Logic02Soal01arr01(n int) {
	//1. create array
	newarray := array2.NewNumberArray(n, n)

	//2. isi array
	ang := 3                             // 2.1 angka awal
	for b := 0; b < len(newarray); b++ { // 2.2 loop isi hingga max array length
		for k := 0; k < len(newarray); k++ { // 2.3 loop kolom hingga max array length
			// 2.4 insert bilangan kedalam (array[baris][kolom])
			newarray[b][k] = int32(ang)
		} // loop isi kolom ends
		ang += 3 // 2.5 setiap b baru, angka = angka +3
	}
	//3 print array yg telah diisi
	array2.PrintNumberArray(newarray)
}

func Logic02Soal01arr02(n int) {
	//1 create array
	array3 := array2.NewNumberArray(n, n)

	//2 isi array
	bil := 3                                       //2.1 inisiasi bilangan awal
	for baris := 0; baris < len(array3); baris++ { // 2.2 loop isi baris
		for kolom := 0; kolom < len(array3); kolom++ { // 2.3 loop isi kolom
			// 2.4 insert bilangan kedalam array2
			array3[baris][kolom] = int32(bil)
		}
		// angka +3 setiap baris baru
		bil += 3
	}
	//3 print array2
	array2.PrintNumberArray(array3)
}

package logic02

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic02Soal02arr00(n int) {
	//1 create array
	array1 := array2.NewNumberArray(n, n)

	//2 isi array menggunakan loop baris dan kolom
	for baris := 0; baris < len(array1); baris++ { // 2.2 loop isi baris : start
		angka := 3                                     // 2.1 inisiasi angka awal & 2.6 setiap baris baru, angka kembali ke = 3
		for kolom := 0; kolom < len(array1); kolom++ { // 2.3 loop isi kolom : start
			//2.4 insert bilangan kedalam array1[baris][kolom]
			array1[baris][kolom] = int32(angka)
			angka += 3 // 2.5 setiap kolom baru, angka = angka + 3
		} // loop isi kolom : ends
	} //loop isi baris : ends
	//3 print array yg telah diisi
	array2.PrintNumberArray(array1)
}

func Logic02Soal02arr01(n int) {
	//1 create array
	array4 := array2.NewNumberArray(n, n)

	//2 isi array with loop
	for row := 0; row < len(array4); row++ { // 2.1 loop row : start
		// 2.3 inisiasi bilangan awal SEKALIGUS setiap baris baru, bilangan kembali ke = 3
		bilangan := 3
		for col := 0; col < len(array4); col++ { // 2.2 loop col : start
			// 2.4 insert bilangan ke dalam array4
			array4[row][col] = int32(bilangan)
			bilangan += 3 // 2.5 setiap kolom baru, angka ditambah +3
		} // 2.2 loop col : end
	} // 2.1 loop row : end
	//3 print array yg telah diisi bilangan
	array2.PrintNumberArray(array4)
}

func Logic02Soal02arr02(n int) {
	//1 create array
	array5 := array2.NewNumberArray(n, n)
	//2 isi array dengan for loop
	for baris := 0; baris < len(array5); baris++ {
		angka := 3
		for kolom := 0; kolom < len(array5); kolom++ {
			array5[baris][kolom] = int32(angka)
			angka += 3
		}
	}
	array2.PrintNumberArray(array5)
}

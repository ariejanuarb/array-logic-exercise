package logic03

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

/*
Soal 03
n = 9
	0	1	2	3	4	5	6	7	8
0		3	3	3	3
1			6	6	6				6
2				9	9			9	9
3					12		12	12	12
4	15	15	15	15		15	15	15	15
5	12	12	12		12
6	9	9			9	9
7	6				6	6	6
8					3	3	3	3
*/

func Logic03Soal03arr00(n int) {
	// 1. create array
	array := array2.NewNumberArray(n, n) // 1. create array

	// 2. isi array
	angka := 3                        // 2.1 angka awal = 3
	for i := 0; i < len(array); i++ { // 2.2 loop baris : start
		nt := n / 2                          // 2.3 nilai tengah sebagai batas turunnya angka
		for j := 0; j < len(array[i]); j++ { // 2.2 loop kolom : start
			/* temukan pola pada koordinat (baris,kolom) dari kotak yang berisi angka*/
			// 2.5 buat 1/2 segitiga kiri atas
			if i < j && j <= nt { //ketika i < j (baris < kolom), dari awal kolom sampai batas tengah (diatas batas tengah tidak ada bilangan)
				array[i][j] = int32(angka) // angka awal(int) di convert ke int32
				// 2.5.1 pantulkan rumus diatas untuk buat 1/2 segitiga kanan bawah
			} else if i > j && j >= nt {
				array[i][j] = int32(angka)
				// 2.6 buat 1/2 segitiga kiri bawah (dasar rumus dari logic03 soal01)
			} else if i+j < n-1 && i >= nt { // i+j = 7 && i mulai dari 4 atau >= nt, kenapa pakai i? karena yg mau di print full hanya baris 4 saja
				array[i][j] = int32(angka)
				// 2.6.1 pantulkan rumus diatas untuk buat 1/2 segitiga kanan atas
			} else if i+j > n-1 && i <= nt { // ganti tanda > jadi <.
				array[i][j] = int32(angka)
			}
		}
		// 2.4 kondisi jika nomor baris kurang dari nilai tengah :
		if i < nt { // angka +3 sampai nilai tengah
			angka += 3
		} else {
			angka -= 3 // diatas batas tengah, angka -3
		}
	}
	// 3. print array
	array2.PrintNumberArrayZeroEmpty(array)
}

func Logic03Soal03arr00Step01(n int) { // membuat 1/2 segitiga kiri atas
	// 1. create array
	array := array2.NewNumberArray(n, n) // 1. create array

	// 2. isi array
	angka := 3                        // 2.1 angka awal = 3
	for i := 0; i < len(array); i++ { // 2.2 loop baris : start
		nt := n / 2                          // 2.3 nilai tengah sebagai batas turunnya angka
		for j := 0; j < len(array[i]); j++ { // 2.2 loop kolom : start
			/* temukan pola pada koordinat (baris,kolom) dari kotak yang berisi angka*/
			// 2.5 buat 1/2 segitiga kiri atas
			if i < j && j <= nt { //ketika i < j (baris < kolom), dari awal kolom sampai batas tengah (diatas batas tengah tidak ada bilangan)
				array[i][j] = int32(angka) // angka awal(int) di convert ke int32
				// 2.5.1 pantulkan rumus diatas untuk buat 1/2 segitiga kanan bawah
			} /*else if i > j && j >= nt {
				array[i][j] = int32(angka)
				// 2.6 buat 1/2 segitiga kiri bawah (dasar rumus dari logic03 soal01)
			} else if i+j < n-1 && i >= nt { // i+j = 7 && i mulai dari 4 atau >= nt, kenapa pakai i? karena yg mau di print full hanya baris 4 saja
				array[i][j] = int32(angka)
				// 2.6.1 pantulkan rumus diatas untuk buat 1/2 segitiga kanan atas
			} else if i+j > n-1 && i <= nt { // ganti tanda > jadi <.
				array[i][j] = int32(angka)
			}*/
		}
		// 2.4 kondisi jika nomor baris kurang dari nilai tengah :
		if i < nt { // angka +3 sampai nilai tengah
			angka += 3
		} else {
			angka -= 3 // diatas batas tengah, angka -3
		}
	}
	// 3. print array
	array2.PrintNumberArrayZeroEmpty(array)
}

func Logic03Soal03arr00Step02(n int) { // +membuat 1/2 segitiga kiri bawah
	//1 create array
	array0300 := array2.NewNumberArray(n, n)
	//2.1 inisiasi variabel dan value
	angka := 3
	tengah := n / 2
	//2 isi array (angka awal 3, +3 tiap baris baru)
	for baris := 0; baris < len(array0300); baris++ {
		for kolom := 0; kolom < len(array0300); kolom++ {
			if kolom > baris && kolom <= tengah { // 1/2 segitiga atas kiri
				array0300[baris][kolom] = int32(angka)
			} else if kolom < baris && kolom >= tengah { // mirror dari segitiga atas kiri
				array0300[baris][kolom] = int32(angka)
			} else if baris+kolom < n-1 && baris >= tengah { // 1/2 segitiga kiri bawah
				array0300[baris][kolom] = int32(angka)
			} else if baris+kolom > n-1 && baris <= tengah { // mirror dari segitiga kiri bawah
				array0300[baris][kolom] = int32(angka)
			}
		}
		if baris < tengah {
			angka += 3
		} else {
			angka -= 3
		}
	}
	//3 print array
	array2.PrintNumberArrayZeroEmpty(array0300)
}

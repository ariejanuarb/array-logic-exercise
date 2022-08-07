package logic04

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic04Soal05Step01(n int) {
	// 1. create array
	kotak := array2.NewNumberArray(n, n)
	// 1.1 initiate first number
	angka := 1
	for a := 1; a < n; a++ { // isi array : mulai
		kotak[n-1-a][0] = int32(angka) // inisiasi angka = 1 di titik awal kotak[2][0]
		angka++                        // angka = angka + 1
	} // isi array : selesai
	// 3. print array
	array2.PrintNumberArray(kotak)
}

func Logic04Soal05Step02(n int) {
	//1. create array
	kotak := array2.NewNumberArray(n, n)
	//1.1 angka pertama
	awal := 1
	// 2. buat pengulangan untuk tiap sisi (1 s.d 4)
	for sisi := 1; sisi <= 4; sisi++ {
		// 1.2 isi array : mulai. buat titik awal mulai dimasukannya angka awal pada kotak[4-1-1][0] = int32(awal)
		for c := 1; c < n; c++ { // loop isi array : start
			kotak[n-1-c][0] = int32(awal)
			awal++
		} // loop isi array : selesai
	}
	// 1.3 print array yang telah diisi
	array2.PrintNumberArray(kotak)
}

func Logic04Soal05arr00(n int) {
	// 1. create array
	kotak := array2.NewNumberArray(n, n)
	// 1.1 tulis angka awal
	angka := 1

	// 2. buat pengulangan untuk tiap sisi (1 s.d 4 == kotak)
	for sisi := 1; sisi <= 4; sisi++ {
		// 1.2 buat loop isi array
		for b := 1; b < n; b++ { //loop isi array mulai, diisi sebanyak 3x pengulangan
			// 3. buat kondisi pengujian untuk titik awal setiap isi serta nilai didalamnya
			// 3.1 pengujian sisi pertama (sisi kiri, diisi dari bawah ke atas)
			if sisi == 1 {
				kotak[n-1-b][0] = int32(angka) // kotak[n-1-b][0] adalah titik awal
			} else if sisi == 2 { // 3.2 pengujian sisi ke-2 (sisi atas, diisi dari kiri ke kanan)
				kotak[0][b] = int32(angka) // kotak[0][b] adalah titik awal
			} else if sisi == 3 { // 3.3 pengujian sisi ke-3 (sisi kanan, diisi dari atas ke bawah)
				kotak[b][n-1] = int32(angka) // kotak[b][n-1] adalah titik awal
			} else if sisi == 4 { // 3.4 pengujian sisi ke-4 (sisi kanan, diisi dari kanan ke kiri)
				kotak[n-1][n-1-b] = int32(angka)
			}
			angka++ // pertambahan angka dilakukan ketika masing2 titik awal telah dimasukkan bilangan pertama
		} //loop isi array selesai
	}
	// 1.3 print array yg telah diisi
	array2.PrintNumberArrayZeroEmpty(kotak)
}

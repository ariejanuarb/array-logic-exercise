package logic02

import (
	"fmt"
)

func Logic02Soal01(n int) {
	a := 3                   // bilangan awal
	for i := 0; i < n; i++ { // loop baris:start
		for j := 0; j < n; j++ { //loop kolom:start
			fmt.Print(a, "\t")
		} // loop kolom:end
		fmt.Println("\n") // baris baru
		a += 3            // tambah nilai a di setiap baris baru
	} // loop baris:end
}

func Logic02Soal02(n int) {
	// loop baris start
	for i := 0; i < n; i++ {
		// loop kolom
		a := 3
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
			a += 3
		}
		// pindah baris baru
		fmt.Println("\n")
		// loop baris ends
	}
}

func Logic02Soal03(n int) {
	// loop baris x9 start
	for i := 0; i < n; i++ {
		// loop kolom x9 start
		a := 3 * n
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
			a -= 3
		} // loop kolom end
		// pindah baris baru setiap loop kolom selesai
		fmt.Println("\n")
	} // loop baris ends
}

func Logic02Soal04(n int) {
	a := 3 * n
	// loop baris
	for i := 0; i < n; i++ {
		//loop kolom
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
		} // kolom selesai
		// baris baru
		fmt.Println("\n")
		// tambah nilai a di baris selanjutnya
		a -= 3
	}
}

func Logic02Soal05(n int) {
	// nilai tengah
	nt := n / 2
	a := 3
	for i := 0; i < n; i++ { // loop baris:start
		for j := 0; j < n; j++ { //loop kolom:start
			fmt.Print(a, "\t")
		} // loop kolom:end
		if i < nt {
			fmt.Println("\n") // baris baru
			a += 3            // tambah nilai a di baris selanjutnya
		} else {
			fmt.Println("\n") // baris baru
			a -= 3            // tambah nilai a di baris selanjutnya
		}
	} // loop baris:end
}

func Logic02Soal06(n int) {
	nt := n / 2
	for i := 0; i < n; i++ { // loop baris:start
		a := 3                   // nilai awal sekaligus nilai reset setiap looping baris baru dimulai
		for j := 0; j < n; j++ { // loop kolom:start
			if j < nt {
				fmt.Print(a, "\t")
				a += 3
			} else {
				fmt.Print(a, "\t")
				a -= 3
			}
		} // loop kolom:end
		fmt.Println("\n") // pindah baris baru
	} // loop baris:end
}

func Logic02Soal07(n int) {
	a := 3
	for i := 0; i < n; i++ { // loop baris:start
		for j := 0; j < n; j++ { //loop kolom:start
			if i >= j { // jika baris lebih besar dan sama dengan kolom
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		} // loop kolom:end
		// baris baru
		fmt.Println("\n")
		// tambah nilai a di baris selanjutnya
		a += 3
	} // loop baris:end
}

func Logic02Soal08(n int) {
	a := 3
	for i := 0; i < n; i++ { // loop baris:start
		for j := 0; j < n; j++ { //loop kolom:start
			if i <= j {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		} // loop kolom:end
		// baris baru
		fmt.Println("\n")
		// tambah nilai a di baris selanjutnya
		a += 3
	} // loop baris:end
}

func Logic02Soal09(n int) {
	for i := 0; i < n; i++ { // loop baris:start
		a := 0                   // reset nilai a = 0
		for j := 0; j < n; j++ { // loop kolom:start
			a += 3          // nilai awal a = 0 + 3 = 3
			if i+j == n-1 { // kondisi batas tengah diagonal
				fmt.Print(a, "\t") // nilai awal a = 3
			} else if i+j >= n-1 { // kondisi area setelah diagonal
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		} // loop kolom:end
		fmt.Println("\n") // pindah baris baru
	} // loop baris:end
}

func Logic02Soal10(n int) {
	for i := 0; i < n; i++ { // loop baris:start
		a := 0
		for j := 0; j < n; j++ { // loop kolom:start
			a += 3
			if i+j <= n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		} // loop kolom:end
		fmt.Println("\n") // pindah baris baru
	} // loop baris:end
}

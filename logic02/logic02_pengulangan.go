package logic02

import (
	"fmt"
)

/*--------------------------------------*/

func Logic02Soal0100(n int) {
	a := 3                   // starting number
	for i := 0; i < n; i++ { // loop baris:start
		for j := 0; j < n; j++ { //loop kolom:start
			fmt.Print(a, "\t") // print a 9x on every column
		} // loop kolom:end
		fmt.Println("\n") // jump into new rows
		a += 3            // add a by 3 on every new rows
	} // loop baris:end
}

func Logic02Soal0101(n int) {
	b := 3                   // bilangan awal
	for i := 0; i < n; i++ { // loop baris :start
		for j := 0; j < n; j++ { // loop kolom:start
			fmt.Print(b, "\t") // print b sebanyak nx
		} // loop kolom :end ketika j sudah < n
		fmt.Println("\n")
		b += 3
	}
}

func Logic02Soal0102(n int) {
	c := 3
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(c, "\t")
		}
		fmt.Println("\n")
		c += 3
	}
}

func Logic02Soal0103(n int) {
	d := 3
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(d, "\t")
		}
		fmt.Println("\n")
		d += 3
	}
}

func Logic02Soal0104(n int) {
	e := 3
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(e, "\t")
		}
		fmt.Println("\n")
		e += 3
	}
}

func Logic02Soal0105(n int) {
	f := 3
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(f, "\t")
		}
		fmt.Println("\n")
		f += 3
	}
}

/*--------------------------------------*/

func Logic02Soal0200(n int) {
	for baris := 0; baris < n; baris++ { // 1. loop baris start
		angka := 3                           // 3. inisiasi angka awal, dan kembali ke-3 ketika loop baris dimulai
		for kolom := 0; kolom < n; kolom++ { // 2. loop kolom start
			fmt.Print(angka, "\t")
			angka += 3
		}
		fmt.Println("\n")
	}
}

/*--------------------------------------*/

func Logic02Soal0300(n int) {
	for baris := 0; baris < n; baris++ {
		bil := n * 3
		for kolom := 0; kolom < n; kolom++ {
			fmt.Print(bil, "\t")
			bil -= 3
		}
		fmt.Println("\n")
	}
}

/*--------------------------------------*/

func Logic02Soal0400(n int) {
	angka := n * 3
	for baris := 0; baris < n; baris++ {
		for kolom := 0; kolom < n; kolom++ {
			fmt.Print(angka, "\t")
		}
		fmt.Println("\n")
		angka -= 3
	}

}

/*--------------------------------------*/

func Logic02Soal0500(n int) {
	value := 3
	nt := n / 2
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			fmt.Print(value, "\t")
		}
		fmt.Println("\n")
		if row < nt {
			value += 3
		} else {
			value -= 3
		}

	}
}

/*--------------------------------------*/

func Logic02Soal0600(n int) {

}

/*--------------------------------------*/

func Logic02Soal0700(n int) {

}

/*--------------------------------------*/

func Logic02Soal0800(n int) {

}

/*--------------------------------------*/

func Logic02Soal0900(n int) {

}

/*--------------------------------------*/

func Logic02Soal1000(n int) {

}

/*--------------------------------------*/
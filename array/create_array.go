package array

/* function NewString array,
parameter baris int, kolom int,
return value array 2 dimensi tipe string [][]string

logic04 terlalu rumit jika diselesaikan dengan cara manual (buat array dengan \t \n)
lebih efisien langsug pakai struktur data array
*/

func NewStringArray(baris int, kolom int) [][]string {
	// make() = membuat slice, dari array 2D string, parameter baris.
	result := make([][]string, baris) // deklarasikan result sebagai slice baris (ulangi lagi materi slice)
	for i := range result {
		result[i] = make([]string, kolom)
	}
	return result
}

func NewNumberArray(baris int, kolom int) [][]int32 {
	result := make([][]int32, baris)
	for i := range result {
		result[i] = make([]int32, kolom)
	}
	return result
}

package array

import "fmt"

func PrintNumberArray(array [][]int32) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ { // array indeks ke-i
			fmt.Print(array[i][j], "\t") // print array indeks ke-i dan ke-j
		}
		fmt.Println("\n")
	}
}

func PrintNumberArrayZeroEmpty(array [][]int32) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if array[i][j] > 0 {
				fmt.Print(array[i][j], "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Println("\n")
	}
}

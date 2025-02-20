package go_slice_fresh

import "fmt"

func Print(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i], "\t")
	}
}

func Print2D(slice [][]int) {
	fmt.Print("[")
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			fmt.Print(slice[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Print("]")
}

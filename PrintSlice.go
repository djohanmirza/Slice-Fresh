package go_slice_fresh

import "fmt"

func Print(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i])
	}
}

func Print2D(slice [][]int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			fmt.Print(slice[i][j])
		}
		fmt.Println()
	}
}

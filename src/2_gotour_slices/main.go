package main

import (
	"golang.org/x/tour/pic"
	"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	fmt.Println("Pic", dx, dy)	
	// Init 2d matrix
	dy_slice := make([][]uint8, dy)		
	for i:=0; i<dx; i++{
		// Create 1D matrix and set to day_slice[i]		
		tmp := make([]uint8, dx) // Right way
		// tmp := []uint8{3,3} // Wrong way
		dy_slice[i] = tmp
	}
	return dy_slice
}

func main() {
	pic.Show(Pic)
}

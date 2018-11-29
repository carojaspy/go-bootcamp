package main

import (
	"fmt"
)

func checkSqrt(temp float64, z float64) bool {
	diff := temp - z
	if diff == 0.0 {
		return true
	}
	return false
}//end CheckSqrt

func Sqrt(x float64) float64 {
	fmt.Println("Calculating Sqrt of: ", x)
	z := 1.0
	temp := 0.0
	// iterate 10 times
	for i:=1.0 ; i<10; i++ {
		temp = z - (z*z - x) / (2*z)
		fmt.Println(i,"Prev Val:", z, "New Val: ", temp)
		// Checking if it's aproximated enough 
		if checkSqrt(temp, z) {
			return temp
		}
		z = temp		
	}
	return z
}// endSqrt

func main() {
	fmt.Println(Sqrt(4))
}


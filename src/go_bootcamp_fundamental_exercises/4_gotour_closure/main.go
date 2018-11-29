package main

import "fmt"

func iterative_fibonacci(fib int) int{
	// fmt.Println("calculating fibbonaci of: ", fib)
	if fib  <= 0{
		fmt.Println("calculating fibbonaci of: ", fib, "is : ", 0)
		return 0
	}
	if fib == 1{
		fmt.Println("calculating fibbonaci of: ", fib, "is : ", 1)
		return 1
	}
	var a = 0
	b := 1
	c := 0
	for i:= 1; i<fib; i++ {
		print(i)
		c = a + b
		a = b
		b = c
	}
	// fmt.Println(a, b, c)
	fmt.Println("calculating fibbonaci of: ", fib, "is : ", c)
	return c
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib := 0
	//Defining func that returns an int
	closure_func := func () int {
		fmt.Println("Closure of fibbonaci of ", fib)
		if fib  <= 0{
			fmt.Println("fibbonaci of: ", fib, "is : ", 0)
			// Increase for next iteration
			fib++
			return 0
		}
		if fib == 1{
			fmt.Println("fibbonaci of: ", fib, "is : ", 1)
			// Increase for next iteration
			fib++
			return 1
		}
		// Increase for next iteration
		fib++
		var a = 0
		b := 1
		c := 0
		for i:= 1; i<fib; i++ {
			print(i)
			c = a + b
			a = b
			b = c
		}
		fmt.Println("fibbonaci of: ", fib, "is : ", c)
		// Return fibonacci
		return c
	}//end closure
	return closure_func
}//End fibonacci


func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
		// fmt.Println(iterative_fibonacci(i))
	}
}


package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x := 0
	x_1 := 1
	result := 0
	return func() int {

		result = x + x_1
		temp := x
		x, x_1 = x_1, result
		return temp

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

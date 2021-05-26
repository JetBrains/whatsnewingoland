package main

import "fmt"

// Step 1. To learn why this syntax is useful, let's bring up the builtin Terminal
//		   Then invoke the following command:
//		   go build -gcflags="-d=ssa/check_bce/debug=1" main.go
//		   This will tell us where the bounds check is performed. Since this slows down
//		   Our code every time it's performed, removing it is a good thing.

func main() {
	s := []int{1, 2}
	a := (*[2]int)(s)

	fmt.Printf("The sumSlice is: %d\n", sumSlice(s))
	fmt.Printf("The sumArray is: %d\n", sumArray(a))
}

func sumSlice(s []int) int {
	return s[0] + s[1]
}

func sumArray(s *[2]int) int {
	return s[0] + s[1]
}

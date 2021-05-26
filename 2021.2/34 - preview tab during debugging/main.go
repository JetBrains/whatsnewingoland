package main

import (
	"fmt"

	"wni20212/previewtabduringdebugging/sum"
)

func main() {
	a := 1
	b := 2
	s := sum.Do(a, b)
	fmt.Printf("The sum of %d + %d = %v", a, b, s)
}

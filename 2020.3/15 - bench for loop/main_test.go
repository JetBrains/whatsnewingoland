package main

import "testing"

func BenchmarkDemo(b *testing.B) {
	d := demoType{Field1: 1}
	val := 0

	// Step 1. Type for and select from the list "for i := 0; i < b.N; i++ {"
	// Step 2. Move the val increment line inside the new for loop
	for i := 0; i < b.N; i++ {
		val = d.AddToField(i)
	}

	b.Logf("val is: %d\n", val)
}

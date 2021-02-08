package main

import "testing"

// Step 1. Replicate the test below by typing"fun"
//		   and selecting "func Test" from the list
func TestDemo(t *testing.T) {
	d := demoType{Field1: 1}
	val := d.AddToField(0)
	if d.Field1 != val {
		t.Fatalf("%d != %d", d.Field1, val)
	}
}

func TestDemoType_AddToField(t *testing.T) {
}

// Step 2. Replicate the benchmark below by typing "fun"
// 		   and selecting "func Bench" from the list
func BenchmarkDemo(b *testing.B) {
	d := demoType{Field1: 1}
	val := 0
	for i := 0; i < b.N; i++ {
		val = d.AddToField(i)
	}
	b.Logf("val is: %d\n", val)
}

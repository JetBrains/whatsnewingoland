package main

// Step 1. Use the "Generate getter and setter" or
//		   the individual "Generate getter" / "Generate setter"
//		   intentions on the struct below.
// Shortcut: Alt + Enter on Windows/Linux
//           ⌥ + Enter on macOS

type demo struct {
	field1 int
	field2 string
	field3 struct {
		field4 bool
	}
}

var _ demo
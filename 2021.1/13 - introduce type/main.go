package main

type demo struct {
	field1 int
	field2 string

	// Step 1. Invoke the "Introduce type" refactoring on the struct below
	// Shortcut: Ctrl + Alt + Shift + T on Windows/Linux
	//           Shift + T on macOS

	field3 struct {
		field4 bool
	}
}

var _ demo

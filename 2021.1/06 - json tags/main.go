package main

// Step 1. Invoke the "Add key to tags" intention to
//		   to add your desired key to all the struct tags.
// Shortcut: Alt + Enter on Windows/Linux
//           ⌥ + Enter on macOS

type Demo struct {
	Field1 *struct {
		Field3 int
	}
	Field2 string
	Field3 bool
}

func _() {
	_ = Demo{}
}

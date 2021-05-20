package main

// Step 1. Invoke the "Add key to tags" intention to
//		   to add your desired key to all the struct tags.
// Shortcut: Alt + Enter on Windows/Linux
//           ⌥ + Enter on macOS

type Demo struct {
	Field1 int
	Field2 string
	Field3 *struct {
		SubField1 int
	}
}

func _() {
	_ = Demo{}
}

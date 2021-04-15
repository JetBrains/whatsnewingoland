package main

func main() {
	// Step 1. Invoke the "Create type" quick-fix on the "demo" struct
	//		   initialization below.
	// Shortcut: Alt + Enter on Windows/Linux
	//           ⌥ + Enter on macOS

	_ = Demo{
		Field1: 1,
		Field2: "two",
		Field3: struct{ Field4 bool }{
			Field4: true,
		},
	}
}

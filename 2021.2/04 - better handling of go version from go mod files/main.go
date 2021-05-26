package main

import (
	_ "embed"
	"fmt"
)

// Step 1. Invoke the Show Context Actions menu and choose
//		   Set Go version in go.mod to 1.16
// Shortcut: Alt + Enter on Windows/Linux
//           ⌥ + Enter on macOS

//go:embed go.mod
var module string

func main() {
	fmt.Println(module)
}

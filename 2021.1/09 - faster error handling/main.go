package main

import (
	"errors"
	"fmt"
)

var ErrInvalidConst = errors.New("could not process string")

func process(a string) (string, error) {
	if a == "Hello!" {
		return "Hello, World!", nil
	}

	return "", ErrInvalidConst
}

func demo(a string) error {
	// Step 1. Invoke the "Handle error" quick-fix on "process()"
	// Shortcut: Alt + Enter on Windows/Linux
	//           ⌥ + Enter on macOS

	process(a)

	fmt.Println(s)
	return nil
}

func main() {
	// Step 2. Invoke the same quick-fix on "demo()"

	demo("Hello!")
}

package main

import "testing"

func TestDemo(t *testing.T) {
	t.Logf("From the TestDemo\n")

	go func() {
		// Step 1. Fatal* calls done in goroutines should be
		// Shortcut: Alt + Enter on Windows/Linux
		//           ⌥ + Enter on macOS

		t.Fatalf("Fatal error from goroutine\n")
	}()
}

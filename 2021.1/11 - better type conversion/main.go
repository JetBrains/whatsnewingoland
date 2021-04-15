package main

func _() {
	// Step 1. Invoke the "Convert to 'string'" quick-fix to fix the issue.
	// Shortcut: Alt + Enter on Windows/Linux
	//           ⌥ + Enter on macOS

	m := []rune{'H', 'e', 'l', 'l', 'o'}
	_ = []string{
		m,
	}
}

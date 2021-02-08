package main

func main() {
	var err error
	a, err := 1, nil // Here "err" is not shadowed, it's reassigned
	println(a, err)

	// Step 1. Observe the color change of "a"
	// Step 2. Invoke Alt + Enter to use the "Navigate to shadowed declaration" intention
	for _, a := range []string{"Hello", ",", "World", "!", "\n"} {
		print(a)
	}
}

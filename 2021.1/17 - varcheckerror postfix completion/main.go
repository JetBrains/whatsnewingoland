package main

import "encoding/json"

func demo() error {
	myValue, result := []byte{}, interface{}(nil)

	// Step 1. Type .vce after Unmarshal()

	json.Unmarshal(myValue, &result)

	_ = result
	return nil
}

func main() {
	// Step 2. Type .vce after Unmarshal()

	demo()
}
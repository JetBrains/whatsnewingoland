package main

import (
	"encoding/asn1"
	"fmt"
)

type demo struct {
	Field1 int
	Field2 string
}

func main() {
	myValue := demo{1, "two"}
	raw, _ := asn1.Marshal(myValue)

	// As with the (json/xml).Unmarshal function call,
	// the IDE now detects the asn1.Unmarshal calls into non-pointer types

	// Step 1. Invoke the "Prepend '&'" quick-fix to resolve the problem
	// Shortcut: Alt + Enter on Windows/Linux
	//           ⌥ + Enter on macOS

	var result demo
	_, _ = asn1.Unmarshal(raw, result)

	fmt.Println(result)
}

// The strings package provides a String type that wraps the built-in Go string type.
//
// # String Type
//
// The string type is a wrapper around the built-in Go string type.
//
// # Example Usage
//
//	package main
//
//	import (
//
//		"fmt"
//
//		"github.com/cramanan/go-types/strings"
//
//	)
//
//	func main() {
//		// Convert from string
//		fromString := From{"Hello World!"}
//
//		fmt.Printf("%q", fromString) // Output: 'Hello World !'
//
//		// Use a method
//		fromString.At(-1) // returns "!"
//	}
package strings

type String string

// Returns a new empty String //
func New() String { return "" }

func From[S ~string | ~[]byte | ~[]rune](value S) String { return String(value) }

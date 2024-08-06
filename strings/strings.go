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

// String is a type that wraps the built-in string type,
// allowing for custom methods and behaviors to be defined.
type String string

// IString is an interface that can be satisfied by types that are assignable to string, []byte, or []rune.
type IString interface {
	~string | ~[]byte | ~[]rune
}

// IChar is an interface that can be satisfied by types that are assignable to byte or rune.
type IChar interface {
	~byte | ~rune
}

// New returns a new, empty String.
func New() String { return "" }

// From converts a value of type S, which can be a string, []byte, or []rune, to a String.
func From[S IString](value S) String { return String(value) }

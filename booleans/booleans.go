// The booleans package provides Bool type that wraps the built-in Go bool type.
//
// # Boolean Type
//
// The Boolean type is a generic wrapper around the built-in Go Boolean type.
// It allows you to work with Boolean in a Object-Oriented way.
//
// # Example Usage
//
//	package main
//
//	import (
//
//		"fmt"
//
//		"github.com/cramanan/go-types/booleans"
//
//	)
//
//	func main() {
//		b := True <=> var b Boolean = True
//
//		// Use functions
//		AND(b, False) // returns False
//
//		// Use methods
//		False.OR(b) // returns True
//
//	}
package booleans

type Boolean bool

const (
	True  Boolean = true
	False Boolean = false
)

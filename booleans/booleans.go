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

// New returns False as a Boolean.
func New() Boolean {
	return *new(Boolean)
}

// From converts b into a Boolean.
func From[B ~bool](b B) Boolean {
	return Boolean(b)
}

// From converts int into to Boolean.
func FromInt(i int) Boolean { return i != 0 }

// ToInt convert b into an integer.
func ToInt(b Boolean) int {
	if b {
		return 1
	}
	return 0
}

// NOT returns the logical negation of b.
func NOT(b Boolean) Boolean { return !b }

// AND returns the logical conjunction of a and b.
func AND(a, b Boolean) Boolean { return a && b }

// NAND returns the logical negation of the conjunction of a and b.
func NAND(a, b Boolean) Boolean { return !(a && b) }

// OR returns the logical disjunction of a and b.
func OR(a, b Boolean) Boolean { return a || b }

// NOR returns the logical negation of the disjunction of a and b.
func NOR(a, b Boolean) Boolean { return !(a || b) }

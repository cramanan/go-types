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
func New() Boolean { return *new(Boolean) }

// IsTruthy performs a Non Zero comparison of infered type T
func IsTruthy[T comparable](b T) bool { return b != *new(T) }

// NOT returns the logical negation of b.
func NOT[B ~bool](boolean B) B { return !boolean }

// AND returns the logical conjunction of a and b.
func AND[B1, B2 ~bool](boolean1 B1, boolean2 B2) B1 { return boolean1 && B1(boolean2) }

// NAND returns the logical negation of the conjunction of a and b.
func NAND[B1, B2 ~bool](boolean1 B1, boolean2 B2) B1 { return !AND(boolean1, boolean2) }

// OR returns the logical disjunction of a and b.
func OR[B1, B2 ~bool](boolean1 B1, boolean2 B2) B1 { return boolean1 || B1(boolean2) }

// NOR returns the logical negation of the disjunction of a and b.
func NOR[B1, B2 ~bool](boolean1 B1, boolean2 B2) B1 { return !OR(boolean1, boolean2) }

func XOR[B1, B2 ~bool](boolean1 B1, boolean2 B2) B1 { return boolean1 != B1(boolean2) }

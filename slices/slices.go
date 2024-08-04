// The slices package provides a generic Slice type that wraps the built-in Go slice type.
//
// # Slice Type
//
// The Slice type is a generic wrapper around the built-in Go slice type. It allows you to work with slices in a type-safe and generic way.
//
// # Type Parameters
//
//   - T : The type of elements in the slice. Can be any type.
//
// # Example Usage
//
//	package main
//
//	import (
//
//		"fmt"
//
//		"github.com/cramanan/go-types/slices"
//
//	)
//
//	func main() {
//		// Create a new Slice of integers
//		intSlice := slices.Slice[int]{1, 2, 3, 4, 5}
//
//		fmt.Println(intSlice) // Output: [1 2 3 4 5]
//
//		// Use a method
//		has3 := intSlice.ContainsFunc(func(i int) bool { return i == 3 })
//
//		fmt.Println(has3) // Output: true
//	}
package slices

type Slice[T any] []T

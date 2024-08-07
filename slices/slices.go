// The slices package provides generic Slice wrapper for the built-in Go slice type and slices functions.
package slices

// Slice is a generic type that wraps a slice of any type T.
// It provides a way to work with slices in a type-safe manner.
//
// Example:
//
//	// Create a Slice of integers
//	intSlice := Slice[int]{1, 2, 3}
//
//	// Create a Slice of strings
//	strSlice := Slice[string]{"a", "b", "c"}
//
//	// Use the Slice as a regular slice
//	fmt.Println(intSlice[0]) // prints 1
//	fmt.Println(strSlice[1]) // prints "b"
type Slice[T any] []T

// New creates a new Slice from the provided values.
func New[T any](values ...T) Slice[T] { return values }

// From creates a Slice from a given slice or array.
func From[S ~[]O, O any](s S) Slice[O] { return Slice[O](s) }

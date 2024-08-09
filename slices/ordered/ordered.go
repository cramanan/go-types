package ordered

import (
	"golang.org/x/exp/constraints"
)

// Ordered is a slice of ordered elements.
//
// The Ordered type is a generic slice that can be used to store elements that implement the constraints.Ordered interface.
// This includes types such as int, float64, string, and others that have a natural ordering.
type Ordered[O constraints.Ordered] []O

// New creates a new Ordered slice from the provided values.
//
// The New function takes a variable number of arguments of type T, which must implement the constraints.Ordered interface.
// It returns a new Ordered slice containing the provided values.
func New[T constraints.Ordered](values ...T) Ordered[T] { return values }

// From creates a new Ordered slice from an existing slice.
//
// The From function takes a slice of type `slice` that is equivalent to `[]T`, where `T` implements the constraints.Ordered interface.
// It returns a new Ordered slice containing the elements of the original slice.
func From[S ~[]T, T constraints.Ordered](s S) Ordered[T] { return Ordered[T](s) }
